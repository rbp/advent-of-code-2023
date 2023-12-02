package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func isNum(c byte) bool {
	return c >= byte('0') && c <= byte('9')
}

func main() {
	f, err := os.Open(os.Args[1])
	panicIfErr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]byte, 2)
		for i := 0; i < len(line); i++ {
			if isNum(line[i]) {
				digits[0] = line[i]
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if isNum(line[i]) {
				digits[1] = line[i]
				break
			}
		}
		num, err := strconv.Atoi(string(digits))
		panicIfErr(err)
		total += num
	}
	fmt.Println(total)
}
