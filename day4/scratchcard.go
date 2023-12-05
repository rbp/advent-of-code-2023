package main

import (
	"advent"
	"fmt"
	"os"
	"regexp"
)

var colon *regexp.Regexp = regexp.MustCompile(`\s*:\s*`)
var pipe *regexp.Regexp = regexp.MustCompile(`\s*\|\s*`)
var ws *regexp.Regexp = regexp.MustCompile(`\s+`)

func parseLine(line string) (map[string]bool, []string) {
	parts := colon.Split(line, 2)
	numbers := pipe.Split(parts[1], 2)
	winning := ws.Split(numbers[0], -1)
	won := ws.Split(numbers[1], -1)

	winningSet := make(map[string]bool, len(winning))
	for _, w := range winning {
		winningSet[w] = true
	}
	return winningSet, won
}

func countWon(winningSet map[string]bool, won []string) int {
	value := 0
	for _, w := range won {
		if winningSet[w] {
			value++
		}
	}
	return value
}

func cardValue(line string) int {
	winningSet, won := parseLine(line)
	nWon := countWon(winningSet, won)
	if nWon <= 1 {
		return nWon
	}
	return 1 << (nWon - 1)
}

func collecCards(lines []string) int {
	nCards := make([]int, len(lines))

	for i, line := range lines {
		// By definition, we always have at least one copy of the cards we're looking at
		nCards[i]++
		winningSet, won := parseLine(line)
		nWon := countWon(winningSet, won)
		for j := 0; j < nWon && j < len(lines); j++ {
			nCards[i+j+1] += nCards[i]
		}
	}
	sum := 0
	for _, n := range nCards {
		sum += n
	}
	return sum
}

func main() {
	lines := advent.Readlines(os.Args[1])
	total := 0
	for _, line := range lines {
		total += cardValue(line)
	}
	fmt.Printf("Part 1: %d\n", total)

	nCards := collecCards(lines)
	fmt.Printf("Part 2: %d\n", nCards)
}
