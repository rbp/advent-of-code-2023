package main

import (
	"advent"
	"fmt"
	"os"
	"slices"
)

const LEFT = 'l'
const RIGHT = 'r'

type point struct {
	x int
	y int
}

type pdiff [2]int

func (p *point) diff(other *point) pdiff {
	return pdiff{p.x - other.x, p.y - other.y}
}

func (p *point) add(diff pdiff) *point {
	return &point{p.x + diff[0], p.y + diff[1]}
}

func otherSide(side byte) byte {
	if side == LEFT {
		return RIGHT
	}
	return LEFT
}

// The characteristics of all pipe types.
// For each pipe, the keys are the possible moves
// (i.e., the possible directions of movement),
// and the values are the coordinate diffs for points
// to the left and right of the pipe, when coming from each key.
// I know, it sounds confusing (the verbose typing syntax doesn't help).
//
// Just use the first one as an example:
// Pipe='|' moves UP (-1, 0) and DOWN (1, 0).
// If you're coming from ABOVE (-1, 0), then to the LEFT is (0, 1) and to the RIGHT is (0, -1)
// If you're coming from BELOW (1, 0), then to the LEFT is (0, -1) and to the RIGHT is (0, 1)
var Pipes map[byte]map[pdiff]map[byte][]pdiff = map[byte]map[pdiff]map[byte][]pdiff{
	// There's an obvious simmetry here that could make this simpler,
	// but that's not the point of this exercise.
	'|': map[pdiff]map[byte][]pdiff{
		{-1, 0}: {
			LEFT:  []pdiff{{0, 1}},
			RIGHT: []pdiff{{0, -1}}},
		{1, 0}: {
			LEFT:  []pdiff{{0, -1}},
			RIGHT: []pdiff{{0, 1}}}},
	'-': map[pdiff]map[byte][]pdiff{
		{0, -1}: {
			LEFT:  []pdiff{{-1, 0}},
			RIGHT: []pdiff{{1, 0}}},
		{0, 1}: {
			LEFT:  []pdiff{{1, 0}},
			RIGHT: []pdiff{{-1, 0}}}},
	'L': map[pdiff]map[byte][]pdiff{
		{-1, 0}: {
			RIGHT: []pdiff{{0, -1}, {1, 0}}},
		{0, 1}: {
			LEFT: []pdiff{{1, 0}, {0, -1}}}},
	'J': map[pdiff]map[byte][]pdiff{
		{-1, 0}: {
			LEFT: []pdiff{{0, 1}, {1, 0}}},
		{0, -1}: {
			RIGHT: []pdiff{{1, 0}, {0, 1}}}},
	'7': map[pdiff]map[byte][]pdiff{
		{0, -1}: {
			LEFT: []pdiff{{-1, 0}, {0, 1}}},
		{1, 0}: {
			RIGHT: []pdiff{{-1, 0}, {0, 1}}}},
	'F': map[pdiff]map[byte][]pdiff{
		{1, 0}: {
			LEFT: []pdiff{{0, -1}, {-1, 0}}},
		{0, 1}: {
			RIGHT: []pdiff{{-1, 0}, {0, -1}}}},
}

func pipeMoves(pipe byte) ([]pdiff, error) {
	if moves, ok := Pipes[pipe]; ok {
		var keys []pdiff
		for pt := range moves {
			keys = append(keys, pt)
		}
		return keys[:2], nil
	}
	return []pdiff{}, fmt.Errorf("No moves for pipe %c", pipe)
}

type maze [][]byte

func mazeFromInput(lines []string) maze {
	maze := make(maze, len(lines))
	for i, line := range lines {
		maze[i] = []byte(line)
	}
	return maze
}

func emptyMaze(template maze) maze {
	newMaze := make(maze, len(template))
	for i := range newMaze {
		newMaze[i] = make([]byte, len(template[i]))
		for j := range newMaze[i] {
			newMaze[i][j] = '.'
		}
	}
	return newMaze
}

func outOfBounds(maze maze, p *point) bool {
	return p.x < 0 || p.x >= len(maze) || p.y < 0 || p.y >= len(maze[0])
}

func firstMoves(maze maze, pos *point) (*point, *point) {
	if c := maze[pos.x][pos.y]; c != 'S' {
		panic(fmt.Sprintf("First move must be from 'S' (got %c)", c))
	}

	// There's a guarantee that there will be exactly two valid pipes to go from S
	valid := map[[2]int][]byte{
		{-1, 0}: {'|', '7', 'F'},
		{1, 0}:  {'|', 'L', 'J'},
		{0, -1}: {'-', 'L', 'F'},
		{0, 1}:  {'-', 'J', '7'},
	}
	var adjacent []*point
	for delta, chars := range valid {
		adj := point{x: pos.x + delta[0], y: pos.y + delta[1]}
		if outOfBounds(maze, &adj) {
			continue
		}
		if c := maze[adj.x][adj.y]; slices.Contains(chars, c) {
			adjacent = append(adjacent, &adj)
			if len(adjacent) == 2 {
				break
			}
		}
	}
	return adjacent[0], adjacent[1]
}

func next(maze maze, pos *point, from *point) *point {
	pipe := maze[pos.x][pos.y]
	var s1, s2 *point
	var moves []pdiff
	var err error
	if moves, err = pipeMoves(pipe); err != nil {
		panic(fmt.Sprintf("Impossible to move: at maze[%v] == %c, from=(%v)\n", pos, pipe, from))
	}
	s1 = &point{x: pos.x + moves[0][0], y: pos.y + moves[0][1]}
	s2 = &point{x: pos.x + moves[1][0], y: pos.y + moves[1][1]}

	if *from == *s1 {
		return s2
	}
	return s1
}

func startingPoint(maze maze) *point {
	for i, line := range maze {
		for j, c := range line {
			if c == 'S' {
				return &point{i, j}
			}
		}
	}
	panic("No starting point found")
}

func mark(maze maze, p *point, marker byte) byte {
	if outOfBounds(maze, p) {
		return 'o'
	}
	if maze[p.x][p.y] == '.' {
		maze[p.x][p.y] = marker
		return marker
	}
	// Hit a pipe
	return 0
}

// traverse follows the pipe path in the maze, starting at 'S', and returns the furthest number of steps
// you can be from the start, in the path.
func traverse(maze maze) (int, maze) {
	// Auxiliary matrix, similar to maze, to store [i]n or [o]ut states later
	// (could be created later, but this saves us an extra traverse)
	inOut := emptyMaze(maze)

	s := startingPoint(maze)
	mark(inOut, s, 'S')

	// The idea is to run the maze in both directions at once, stop when they meet and count the steps
	side1, side2 := firstMoves(maze, s)
	var visited []point
	var steps int
	side1From := s
	side2From := s
	for steps = 0; !slices.Contains(visited, *side1) && !slices.Contains(visited, *side2); steps++ {
		visited = append(visited, *side1, *side2)
		mark(inOut, side1, maze[side1.x][side1.y])
		mark(inOut, side2, maze[side2.x][side2.y])

		side1, side1From = next(maze, side1, side1From), side1
		side2, side2From = next(maze, side2, side2From), side2
	}
	return steps, inOut
}

// Return a map of LEFT/RIGHT counts, and IN (empty, or inferred as LEFT or RIGHT)
func markAround(maze maze, p *point, from *point) (map[byte]int, byte) {
	var IN byte
	count := map[byte]int{
		LEFT:  0,
		RIGHT: 0,
	}

	pipe := maze[p.x][p.y]
	var sides map[byte][]pdiff
	var ok bool
	if sides, ok = Pipes[pipe][from.diff(p)]; !ok {
		panic(fmt.Sprintf("Invalid pipe characteristic: pipe=%c, p=%v, from=%v\n", pipe, p, from))
	}
	for _, iside := range []byte{LEFT, RIGHT} {
		side := byte(iside)
		if diffs, ok := sides[side]; ok {
			for _, diff := range diffs {
				switch b := mark(maze, p.add(diff), side); b {
				case 'o':
					IN = otherSide(side)
				case LEFT, RIGHT:
					count[side]++
				}
			}
		}
	}
	return count, IN
}

func countIn(maze maze) int {
	// The idea is to traverse the maze in an arbirary direction, marking each side (i.e., not the direction of movement) LEFT and RIGHT.
	// Counting each as we go.
	// If one side is a wall, that side is OUT and the other one is IN.
	// THEN: Sweep the board from top to bottom. Any '.' touching LEFT is LEFT, any touching RIGHT is RIGHT.
	// If there are '.' touching both LEFT and RIGHT, that's an error.
	// Return the count of the IN.

	s := startingPoint(maze)
	// TODO: I need to mark around S as well
	nextPoint, _ := firstMoves(maze, s)
	pointFrom := s
	var thisCount map[byte]int
	count := map[byte]int{
		LEFT:  0,
		RIGHT: 0,
	}
	var IN, OUT, thisIn byte

	for *nextPoint != *s {
		thisCount, thisIn = markAround(maze, nextPoint, pointFrom)
		if thisIn != 0 {
			if IN != 0 && IN != thisIn {
				panic("IN is not consistent")
			}
			IN = thisIn
		}
		count[LEFT] += thisCount[LEFT]
		count[RIGHT] += thisCount[RIGHT]

		nextPoint, pointFrom = next(maze, nextPoint, pointFrom), nextPoint
	}
	// By now, shall we assume we know IN?
	if IN == LEFT {
		OUT = RIGHT
	} else if IN == RIGHT {
		OUT = LEFT
	} else {
		panic("IN is still 0")
	}

	// Now, for the sweep
	prev := make([]byte, len(maze[0]))
	for i := range prev {
		prev[i] = OUT
	}
	for _, row := range maze {
		for j := range row {
			if row[j] == '.' {
				row[j] = prev[j]
				if slices.Contains([]byte{LEFT, RIGHT}, row[j]) {
					count[row[j]]++
				}
			} else if slices.Contains([]byte{LEFT, RIGHT}, row[j]) {
				prev[j] = row[j]

			}
		}
	}

	fmt.Printf("Final Count (IN=%v) is %v\n", string(IN), count)
	return count[IN]
}
func printBoard(maze maze) {
	for _, row := range maze {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func main() {
	maze := mazeFromInput(advent.Readlines(os.Args[1]))
	dist, inOut := traverse(maze)
	c := countIn(inOut)
	// printBoard(inOut)
	fmt.Printf("Part 1: %d\n", dist)
	fmt.Printf("Part 2: %d\n", c)
}
