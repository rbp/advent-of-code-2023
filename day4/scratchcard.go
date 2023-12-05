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

func cardValue(line string) int {
	parts := colon.Split(line, 2)
	numbers := pipe.Split(parts[1], 2)
	winning := ws.Split(numbers[0], -1)
	won := ws.Split(numbers[1], -1)

	winningSet := make(map[string]bool, len(winning))
	for _, w := range winning {
		winningSet[w] = true
	}

	value := 0
	for _, w := range won {
		if winningSet[w] {
			if value == 0 {
				value = 1
			} else {
				value <<= 1
			}
		}
	}
	return value
}

func main() {
	total := 0
	for _, line := range advent.Readlines(os.Args[1]) {
		total += cardValue(line)
	}
	fmt.Printf("Part 1: %d\n", total)
}
