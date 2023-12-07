package advent

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

// PanicIfErr panics is err is not nil
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Readlines reads filename and returns its lines, one by one
func Readlines(filename string) []string {
	// A more robust way would be a scanner:
	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	...
	// but for the purposes of the advent, we can assume everything fits in memory
	bytes, err := os.ReadFile(filename)
	PanicIfErr(err)
	return strings.Split(strings.TrimRight(string(bytes), "\n"), "\n")
}

// IsNum returns true if c is a (byte represetation of a) number
func IsNum(c byte) bool {
	return c >= byte('0') && c <= byte('9')
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	PanicIfErr(err)
	return n
}

func LineToNumbers(line string) []int {
	stringed := regexp.MustCompile(`\s+`).Split(line, -1)
	nums := make([]int, len(stringed))
	for i, s := range stringed {
		nums[i] = MustAtoi(s)
	}
	return nums
}
