package main

import (
	"advent"
	"fmt"
	"os"
)

func allZeroes(nums []int) bool {
	for _, i := range nums {
		if i != 0 {
			return false
		}
	}
	return true
}

// tillAllZeroes runs through the steps until it's all zeroes
// returns 2 lists: first and last numbers of each iteration
func tillAllZeroes(steps []int) ([]int, []int) {
	first := []int{steps[0]}
	last := []int{steps[len(steps)-1]}
	for it := 0; !allZeroes(steps[:len(steps)-it]); it++ {
		for i := 0; i < len(steps)-1-it; i++ {
			steps[i] = steps[i+1] - steps[i]
		}
		first = append(first, steps[0])
		last = append(last, steps[len(steps)-1-it-1])
	}
	return first, last

}

func prediction(line string) (int, int) {
	nums := advent.LineToNumbers(line)
	first, last := tillAllZeroes(nums)

	var predLast int
	for _, i := range last {
		predLast += i
	}

	var predFirst int
	for i := len(first) - 1; i > 0; i-- {
		predFirst = first[i-1] - predFirst
	}

	return predFirst, predLast
}

func sumPredictions(lines []string) (int, int) {
	var sumFirst, sumLast int
	for _, line := range lines {
		first, last := prediction(line)
		sumFirst += first
		sumLast += last
	}
	return sumFirst, sumLast

}

func main() {
	lines := advent.Readlines(os.Args[1])
	first, last := sumPredictions(lines)
	fmt.Printf("Part 1: %d\n", last)
	fmt.Printf("Part 2: %d\n", first)
}
