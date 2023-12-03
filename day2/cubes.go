package main

import (
	"advent"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reDraw *regexp.Regexp = regexp.MustCompile(`(?P<qty>\d+) (?P<colour>red|blue|green)(?:, |$)?`)
var reGameID *regexp.Regexp = regexp.MustCompile(`(?:^Game )(\d+)`)

func isValid(draw map[string]int) bool {
	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for colour, qty := range draw {
		if max[colour] < qty {
			return false
		}
	}
	return true
}

func parseDraw(s string) map[string]int {
	draw := make(map[string]int)

	for _, match := range reDraw.FindAllStringSubmatch(s, -1) {
		qty, err := strconv.Atoi(match[1])
		advent.PanicIfErr(err)
		draw[match[2]] = qty
	}
	return draw
}

// parseLine returns a gameID and a list of draws
func parseLine(line string) (int, []map[string]int) {
	parts := strings.Split(line, ":")
	idPart := strings.TrimSpace(parts[0])
	drawsPart := strings.TrimSpace(parts[1])

	match := reGameID.FindStringSubmatch(idPart)
	gameID, err := strconv.Atoi(match[1])
	advent.PanicIfErr(err)

	lineDraws := strings.Split(drawsPart, ";")
	draws := make([]map[string]int, len(lineDraws))
	for _, drawStr := range lineDraws {
		draw := parseDraw(strings.TrimSpace(drawStr))
		draws = append(draws, draw)
	}
	return gameID, draws

}

func gameIDIfValid(line string) int {
	gameID, draws := parseLine(line)

	for _, draw := range draws {
		if !isValid(draw) {
			return 0
		}
	}
	return gameID
}

func gamePower(line string) int {
	_, draws := parseLine(line)

	min := make(map[string]int)
	for _, draw := range draws {
		for colour, qty := range draw {
			if min[colour] < qty {
				min[colour] = qty
			}
		}
	}
	return min["red"] * min["green"] * min["blue"]
}

func main() {
	total := 0
	totalPower := 0
	for _, line := range advent.Readlines(os.Args[1]) {
		total += gameIDIfValid(line)
		totalPower += gamePower(line)
	}
	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", totalPower)
}
