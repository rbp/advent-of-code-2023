package main

import (
	"advent"
	"fmt"
	"os"
	"strconv"
)

func isNum(c byte) bool {
	return c >= byte('0') && c <= byte('9')
}

func isSpelledOut(s string) (byte, error) {
	// Returns the number spelled out in the string, as a byte
	// The spelled out number must start at the start of the string, but may have a suffix
	// If the string does not spell out a number, returns an error.
	trie := newTrieNode()
	numbers := map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	for literal, num := range numbers {
		trie.insert(literal)
		if trie.findSubstr(s) {
			return num, nil
		}
	}
	return 0, fmt.Errorf("Could not find a number in %s", s)
}

func findNumberAt(line string, pos int) (byte, error) {
	if isNum(line[pos]) {
		return line[pos], nil
	}
	if n, err := isSpelledOut(line[pos:]); err == nil {
		return n, nil
	}
	return 0, fmt.Errorf("No number in line[%d]", pos)

}

func findFirstNum(line string) byte {
	for i := 0; i < len(line); i++ {
		if n, err := findNumberAt(line, i); err == nil {
			return n
		}
	}
	return 0
}

func findLastNum(line string) byte {
	for i := len(line) - 1; i >= 0; i-- {
		if n, err := findNumberAt(line, i); err == nil {
			return n
		}
	}
	return 0
}

func lineValue(line string) int {
	digits := []byte{
		findFirstNum(line),
		findLastNum(line),
	}
	num, err := strconv.Atoi(string(digits))
	advent.PanicIfErr(err)
	return num

}

func main() {
	total := 0
	for _, line := range advent.Readlines(os.Args[1]) {
		total += lineValue(line)
	}
	fmt.Println(total)
}
