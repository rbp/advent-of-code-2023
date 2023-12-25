package main

import (
	"advent"
	"fmt"
	"os"
	"slices"
)

type point struct {
	x int
	y int
}

var Moves map[byte][2][2]int = map[byte][2][2]int{
	'|': {{-1, 0}, {1, 0}},
	'-': {{0, -1}, {0, 1}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'F': {{0, 1}, {1, 0}},
}

func firstMoves(maze []string, pos *point) (*point, *point) {
	if c := maze[pos.x][pos.y]; c != 'S' {
		panic(fmt.Sprintf("First move must be from 'S' (got %c)", c))
	}
	valid := map[[2]int][]byte{
		{-1, 0}: {'|', '7', 'F'},
		{1, 0}:  {'|', 'L', 'J'},
		{0, -1}: {'-', 'L', 'F'},
		{0, 1}:  {'-', 'J', '7'},
	}
	var adjascent []*point
	for delta, chars := range valid {
		adj := point{x: pos.x + delta[0], y: pos.y + delta[1]}
		if adj.x < 0 || adj.x > len(maze) || adj.y < 0 || adj.y > len(maze[0]) {
			continue
		}
		if c := maze[adj.x][adj.y]; slices.Contains(chars, c) {
			adjascent = append(adjascent, &adj)
			if len(adjascent) == 2 {
				break
			}
		}
	}
	return adjascent[0], adjascent[1]
}

func next(maze []string, pos *point, from *point) *point {
	pipe := maze[pos.x][pos.y]
	var s1, s2 *point
	var moves [2][2]int
	var ok bool
	if moves, ok = Moves[pipe]; !ok {
		panic(fmt.Sprintf("Impossible to move: maze[%v] == %c\n", pos, pipe))
	}
	s1 = &point{x: pos.x + moves[0][0], y: pos.y + moves[0][1]}
	s2 = &point{x: pos.x + moves[1][0], y: pos.y + moves[1][1]}

	if *from == *s1 {
		return s2
	}
	return s1

}

func longestDistance(maze []string) int {
	var s *point
	for i, line := range maze {
		for j, c := range line {
			if c == 'S' {
				s = &point{i, j}
				break
			}
		}
	}

	// The idea is to run the maze in both directions at once, stop when they meet and count the steps
	side1, side2 := firstMoves(maze, s)
	var visited []point
	var steps int
	side1From := s
	side2From := s
	for steps = 0; !slices.Contains(visited, *side1) && !slices.Contains(visited, *side2); steps++ {
		visited = append(visited, *side1, *side2)
		side1, side1From = next(maze, side1, side1From), side1
		side2, side2From = next(maze, side2, side2From), side2
	}

	return steps
}

func main() {
	maze := advent.Readlines(os.Args[1])
	dist := longestDistance(maze)
	fmt.Printf("Part 1: %d\n", dist)
}
