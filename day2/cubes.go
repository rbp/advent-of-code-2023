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

func lineValue(line string) int {
	parts := strings.Split(line, ":")
	idPart := strings.TrimSpace(parts[0])
	drawsPart := strings.TrimSpace(parts[1])

	match := reGameID.FindStringSubmatch(idPart)
	gameID, err := strconv.Atoi(match[1])
	advent.PanicIfErr(err)

	draws := strings.Split(drawsPart, ";")
	for _, drawStr := range draws {
		draw := parseDraw(strings.TrimSpace(drawStr))
		if !isValid(draw) {
			return 0
		}
	}
	return gameID
}

func main() {
	total := 0
	for _, line := range advent.Readlines(os.Args[1]) {
		total += lineValue(line)
	}
	fmt.Println(total)
}
