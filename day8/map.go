package main

import (
	"advent"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Map struct {
	directions string
	nodes      map[string][2]string

	_nextTurn int
}

func (m *Map) nextTurn() int {
	i := m._nextTurn
	if i == len(m.directions)-1 {
		m._nextTurn = 0
	} else {
		m._nextTurn++
	}
	lrToInt := map[byte]int{
		'L': 0,
		'R': 1,
	}
	return lrToInt[m.directions[i]]
}

func parseLines(lines []string) *Map {
	m := &Map{directions: lines[0], nodes: make(map[string][2]string)}

	// Line 1 is empty
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " = ")
		reNodes := regexp.MustCompile(`\(([A-Z]{3}), ([A-Z]{3})\)`)
		match := reNodes.FindStringSubmatch(parts[1])
		m.nodes[parts[0]] = [2]string{match[1], match[2]}
	}
	return m
}

func stepsToZZZ(m *Map) int {
	steps := 0
	node := "AAA"
	for node != "ZZZ" {
		node = m.nodes[node][m.nextTurn()]
		steps++
	}
	return steps
}

func main() {
	lines := advent.Readlines(os.Args[1])
	m := parseLines(lines)
	fmt.Printf("Part 1: %d\n", stepsToZZZ(m))

}
