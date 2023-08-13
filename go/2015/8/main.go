package main

import (
	"fmt"
	"strings"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

func countMemChars(str string) (count int) {
	buf := ""
	count = 0

	for _, char := range str {
		if len(buf) == 0 && char != '\\' {
			count++
			continue
		}

		buf = buf + string(char)

		switch true {
		case strings.HasPrefix(buf, `\x`) && len(buf) == 4:
			fallthrough
		case strings.HasPrefix(buf, `\\`) || strings.HasPrefix(buf, `\"`):
			buf = ""
			count++
		}
	}

	return count - 2
}

func stringEncode(str string) string {
	buf := `"`

	for _, char := range str {
		switch char {
		case '\\':
			buf += `\\`
		case '"':
			buf += `\"`
		default:
			buf += string(char)
		}
	}

	buf += `"`

	return buf
}

func main() {
	input := util.ExitIfError(util.LoadInput(inputFile))

	list := strings.Split(string(input), "\n")

	var p1 int
	var p2 int
	for _, str := range list {
		// we are only dealing with ascii so len() will suffice for literal character count
		p1 += len(str) - countMemChars(str)
		p2 += len(stringEncode(str)) - len(str)
	}

	fmt.Printf("(Part 1) Count in code minus count in memory comes out to %v\n", p1)
	fmt.Printf("(Part 2) String Encoded character total minus count in code comes out to  %v\n", p2)
}
