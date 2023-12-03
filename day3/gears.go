package main

import (
	"advent"
	"fmt"
	"os"
	"slices"
	"strings"
)

type number struct {
	value int
	start int
	end   int
}

type symbol = int

type line struct {
	numbers []number
	symbols []symbol
}

func parseLine(s string) line {
	i := 0
	l := line{}
	for ; i < len(s); i++ {
		if s[i] == '.' {
			continue
		}
		if advent.IsNum(s[i]) {
			j := i + 1
			for ; j < len(s) && advent.IsNum(s[j]); j++ {
			}
			l.numbers = append(l.numbers, number{value: advent.MustAtoi(s[i:j]), start: i, end: j - 1})
			i = j - 1
			continue
		}
		l.symbols = append(l.symbols, symbol(i))
		slices.Sort(l.symbols)
	}
	return l
}

func symbolFoundAt(symbols []symbol, idx int) bool {
	_, found := slices.BinarySearch(symbols, idx)
	return found
}

func parseSchematic(s string) int {
	last3Lines := [3]line{
		line{},
		line{},
		line{},
	}

	lines := strings.Split(s, "\n")
	// Add an empty line at the end to make sure the last line is parsed
	lines = append(lines, "")
	sum := 0
	for _, l := range lines {
		last3Lines[0] = last3Lines[1]
		last3Lines[1] = last3Lines[2]
		last3Lines[2] = parseLine(l)

		// Compute the middle line
		// For a number to be added, it must have an adjacent symbol
		midline := last3Lines[1]
		for _, number := range midline.numbers {
			// First, any symbols immediately before or after the number
			for _, n := range [2]int{number.start - 1, number.end + 1} {
				if symbolFoundAt(midline.symbols, n) {
					sum += number.value
					break
				}
			}
			for _, surroundingLine := range [2]line{last3Lines[0], last3Lines[2]} {
				for n := number.start - 1; n <= number.end+1; n++ {
					if symbolFoundAt(surroundingLine.symbols, n) {
						sum += number.value
						break
					}
				}
			}
		}
	}
	return sum
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	advent.PanicIfErr(err)
	n := parseSchematic(strings.TrimSpace(string(bytes)))
	fmt.Printf("Part 1: %d\n", n)
}
