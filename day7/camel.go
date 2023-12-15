package main

import (
	"advent"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var Cards string = "23456789TJQKA"

type turn struct {
	h   *hand
	bid int
}

type hand struct {
	cards string
	// freq -> [cards that show up this many times]
	cardByFreq map[int][]rune
	_type      int
}

func NewTurn(cards string, bid int) *turn {
	t := &turn{h: NewHand(cards), bid: bid}
	return t
}

func NewHand(cards string) *hand {
	h := &hand{cards: cards}
	cardFreq := make(map[rune]int)
	for _, c := range cards {
		cardFreq[c]++
	}
	byFreq := make(map[int][]rune)
	for c, f := range cardFreq {
		byFreq[f] = append(byFreq[f], c)
	}
	h.cardByFreq = byFreq
	h._type = findType(h)
	return h

}

func findType(h *hand) int {
	if _, ok := h.cardByFreq[5]; ok {
		return FiveOfAKind
	}
	if _, ok := h.cardByFreq[4]; ok {
		return FourOfAKind
	}
	if _, ok := h.cardByFreq[3]; ok {
		if _, ok := h.cardByFreq[2]; ok {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if pairs, ok := h.cardByFreq[2]; ok {
		if len(pairs) == 2 {
			return TwoPair
		}
		return OnePair
	}
	return HighCard
}

func cmpHands(h1, h2 *hand) int {
	if h1._type != h2._type {
		return h1._type - h2._type
	}
	for i := 0; i < 5; i++ {
		if h1.cards[i] != h2.cards[i] {
			return strings.Index(Cards, string(h1.cards[i])) - strings.Index(Cards, string(h2.cards[i]))
		}
	}
	return 0

}

func cmpTurns(t1, t2 *turn) int {
	return cmpHands(t1.h, t2.h)
}

func totalWinnings(lines []string) int {
	turns := make([]*turn, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		turns[i] = NewTurn(parts[0], advent.MustAtoi(parts[1]))
	}
	slices.SortFunc(turns, cmpTurns)

	var winnings int
	for i := 0; i < len(turns); i++ {
		winnings += (i + 1) * turns[i].bid
	}
	return winnings
}

func main() {
	lines := advent.Readlines(os.Args[1])
	winnings := totalWinnings(lines)
	fmt.Printf("Part 1: %d\n", winnings)
}
