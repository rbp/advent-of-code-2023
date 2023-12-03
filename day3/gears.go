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

func (n *number) isAdjacentTo(pos int) bool {
	return n.start-1 <= pos && pos <= n.end+1
}

type symbol = int

type gear struct {
	pos      int
	adjacent []*number
}

func (g *gear) isValid() bool {
	return len(g.adjacent) == 2
}

func (g *gear) ratio() int {
	return g.adjacent[0].value * g.adjacent[1].value
}

type line struct {
	numbers []*number
	symbols []symbol
	gears   []*gear
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
			l.numbers = append(l.numbers, &number{value: advent.MustAtoi(s[i:j]), start: i, end: j - 1})
			i = j - 1
			continue
		}
		if s[i] == '*' {
			l.gears = append(l.gears, &gear{pos: i})
			// Do not "continue", because gears [part 2] are also symbols [part 1]
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

// Parses schematic described by s and returns (sum of part numbers [part 1], sum of gear ratios [part 2])
func parseSchematic(s string) (int, int) {
	last3Lines := [3]line{
		line{},
		line{},
		line{},
	}

	lines := strings.Split(s, "\n")
	// Add an empty line at the end to make sure the last line is parsed
	lines = append(lines, "")
	sum := 0
	ratioSum := 0
	for _, l := range lines {
		// Shift lines "down" by one place, and parse the current one
		last3Lines[0] = last3Lines[1]
		last3Lines[1] = last3Lines[2]
		last3Lines[2] = parseLine(l)

		// We're always focusing on computing the middle line
		// For a part number to be added, it must have an adjacent symbol
		midline := last3Lines[1]
		for _, number := range midline.numbers {
			// First, any symbols immediately before or after the number, in the same line
			for _, n := range [2]int{number.start - 1, number.end + 1} {
				if symbolFoundAt(midline.symbols, n) {
					sum += number.value
					break
				}
			}
			// Then, any adjacent symbols in the surrounding lines
			for _, surroundingLine := range [2]line{last3Lines[0], last3Lines[2]} {
				for n := number.start - 1; n <= number.end+1; n++ {
					if symbolFoundAt(surroundingLine.symbols, n) {
						sum += number.value
						break
					}
				}
			}

			// Finally, try to find adjacent gears in all surrounding lines
			for _, surroundingLine := range [3]line{last3Lines[0], last3Lines[1], last3Lines[2]} {
				for _, gear := range surroundingLine.gears {
					if number.isAdjacentTo(gear.pos) {
						gear.adjacent = append(gear.adjacent, number)
						break
					}
				}
			}
		}

		// Then, see if any gears in the 1st line is valid
		// (gears in the first line - the one about to be shifted out -
		// have all been checked for adjacent numbers by now)
		for _, gear := range last3Lines[0].gears {
			if gear.isValid() {
				ratioSum += gear.ratio()
			}
		}

	}
	return sum, ratioSum
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	advent.PanicIfErr(err)
	sum, rationSum := parseSchematic(strings.TrimSpace(string(bytes)))
	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", rationSum)
}
