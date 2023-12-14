package main

import (
	"advent"
	"fmt"
	"math"
	"os"
	"strings"
)

// Time, Distance
type race [2]int

func parseInput(lines []string) []race {
	times := advent.LineToNumbers(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	distances := advent.LineToNumbers(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	if len(times) != len(distances) {
		panic("times and distances must be the same length")
	}

	races := make([]race, len(times))
	for i := 0; i < len(times); i++ {
		races[i] = race{times[i], distances[i]}
	}
	return races
}

// / btnBoundaries return 2 ints (lower, higher) around which you can hold the button and win the race.
// / The boundaries are not supposed to be *tight* (i.e., exactly the boundaries of button
// / press at which you can win the race), but should be guaranteed to *contain* the right values
func btnBoundaries(r race) (int, int) {
	// This is derived from:
	// btn * (time - bnt) > distance [since the time pushing the button is the boat speed]
	// ...
	// btn ~ time ± sqrt(time^2 - 4*distance)/2
	time := float64(r[0])
	dist := float64(r[1])
	lower := math.Floor(
		(time - math.Ceil(math.Sqrt(time*time-4*dist))) / 2)
	higher := math.Floor(
		(time + math.Ceil(math.Sqrt(time*time-4*dist))) / 2)
	return int(lower), int(higher)

}

func bntMinMax(r race) (int, int) {
	low, high := btnBoundaries(r)
	var btn, min, max int
	for btn = low; btn <= high; btn++ {
		if btn*(r[0]-btn) > r[1] {
			break
		}
	}
	min = btn

	for btn = high; btn >= low; btn-- {
		if btn*(r[0]-btn) > r[1] {
			break
		}
	}
	max = btn
	return min, max

}

func waysToWin(r race) int {
	min, max := bntMinMax(r)
	return max - min + 1
}

func allWaysToWin(races []race) int {
	total := 1
	for _, r := range races {
		total *= waysToWin(r)
	}
	return total
}

func main() {
	lines := advent.Readlines(os.Args[1])
	races := parseInput(lines)
	fmt.Printf("Part 1: %v\n", allWaysToWin(races))
}
