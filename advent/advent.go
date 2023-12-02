package advent

import (
	"os"
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
