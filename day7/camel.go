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

func cardValue(c rune, jokersEnabled bool) int {
	if jokersEnabled {
		return strings.Index("J23456789TQKA", string(c))
	}
	return strings.Index("23456789TJQKA", string(c))
}

type turn struct {
	h   *hand
	bid int
}

type hand struct {
	cards string
	// map[freq] -> [cards that show up this many times]
	cardByFreq    map[int][]rune
	_type         int
	jokersEnabled bool
	nJokers       int
}

func NewTurn(cards string, bid int, jokersEnabled bool) *turn {
	t := &turn{h: NewHand(cards, jokersEnabled), bid: bid}
	return t
}

func NewHand(cards string, jokersEnabled bool) *hand {
	h := &hand{cards: cards, jokersEnabled: jokersEnabled}
	cardFreq := make(map[rune]int)
	for _, c := range cards {
		if jokersEnabled && c == 'J' {
			h.nJokers++
		} else {
			cardFreq[c]++
		}
	}
	byFreq := make(map[int][]rune)
	for c, f := range cardFreq {
		byFreq[f] = append(byFreq[f], c)
	}
	h.cardByFreq = byFreq
	h._type = findType(h)
	return h

}

// What a horrible functions...
func findType(h *hand) int {
	if _, ok := h.cardByFreq[5]; ok {
		return FiveOfAKind
	}
	if _, ok := h.cardByFreq[4]; ok {
		return FourOfAKind + h.nJokers
	}
	if _, ok := h.cardByFreq[3]; ok {
		if h.nJokers > 0 {
			return ThreeOfAKind + 1 + h.nJokers
		}
		if _, ok := h.cardByFreq[2]; ok {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if pairs, ok := h.cardByFreq[2]; ok {
		if len(pairs) == 2 {
			if h.nJokers > 0 {
				return FullHouse
			}
			return TwoPair
		}
		switch h.nJokers {
		case 0:
			return OnePair
		case 1:
			return ThreeOfAKind
		case 2:
			return FourOfAKind
		case 3:
			return FiveOfAKind
		}
	}
	switch h.nJokers {
	case 0:
		return HighCard
	case 1:
		return OnePair
	case 2:
		return ThreeOfAKind
	case 3:
		return FourOfAKind
	case 4:
		return FiveOfAKind
	case 5:
		return FiveOfAKind
	}
	panic(("This should not be reached"))

}

func cmpHands(h1, h2 *hand) int {
	if h1._type != h2._type {
		return h1._type - h2._type
	}
	for i := 0; i < 5; i++ {
		if h1.cards[i] != h2.cards[i] {
			return cardValue(rune(h1.cards[i]), h1.jokersEnabled) - cardValue(rune(h2.cards[i]), h2.jokersEnabled)
		}
	}
	panic("Two hands are the same!")
}

func cmpTurns(t1, t2 *turn) int {
	return cmpHands(t1.h, t2.h)
}

func parseTurns(lines []string, jokersEnabled bool) []*turn {
	turns := make([]*turn, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		turns[i] = NewTurn(parts[0], advent.MustAtoi(parts[1]), jokersEnabled)
	}
	return turns
}

func totalWinnings(lines []string, jokersEnabled bool) int {
	turns := parseTurns(lines, jokersEnabled)
	slices.SortFunc(turns, cmpTurns)

	var winnings int
	for i := 0; i < len(turns); i++ {
		winnings += (i + 1) * turns[i].bid
	}
	return winnings
}

func main() {
	lines := advent.Readlines(os.Args[1])
	winnings := totalWinnings(lines, false)
	fmt.Printf("Part 1: %d\n", winnings)

	jWinnings := totalWinnings(lines, true)
	fmt.Printf("Part 2: %d\n", jWinnings)

}
