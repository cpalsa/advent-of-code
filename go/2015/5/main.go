package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/cpalsa/advent-of-code/go/2015/util"
)

const inputFile = "input.txt"

func main() {
	input, err := util.LoadInput(inputFile)
	util.HandleFatal(err)

	list := strings.Split(string(input), "\n")
	count := 0
	for _, str := range list {
		if isNice1(str) {
			count++
		}
	}
	fmt.Printf("(Part 1) Santa has %v nice strings using the first set of rules\n", count)

	count = 0
	for _, str := range list {
		if isNice2(str) {
			count++
		}
	}
	fmt.Printf("(Part 2) Santa has %v nice strings using the second set of rules\n", count)
}

func isNice2(str string) (nice bool) {
	// need two pairs (2 x 2) of runes not overlapping
	if utf8.RuneCountInString(str) < 4 {
		return false
	}

	hasTwoPairs := false
	hasSymmetry := false

	pairs := make(map[string]bool)
	buf := []rune(str)

	for len(buf) > 1 {
		// check the current pair of tokens
		if len(buf) >= 2 {
			slice := buf[:2]

			pair := string(slice[0]) + string(slice[1])

			// this pair already exist?
			if pairs[pair] {
				hasTwoPairs = true
			} else {
				pairs[pair] = true
			}
		}

		// look ahead one from the pair
		if len(buf) >= 3 {
			slice := buf[:3]

			if slice[0] == slice[2] {
				hasSymmetry = true
			}

			// if overlapping with the next token would register another pair we consume 2 tokens
			if slice[0] == slice[1] && slice[1] == slice[2] {
				buf = buf[2:]
				continue
			}
		}

		if hasSymmetry && hasTwoPairs {
			nice = true
			break
		}

		// consume 1 token
		buf = buf[1:]
	}

	return nice
}

func isNice1(str string) (nice bool) {
	// need at least 3 vowels
	if utf8.RuneCountInString(str) < 3 {
		return false
	}

	vowels := 0
	hasDouble := false

	var last rune
	for _, char := range str {
		if isBadString(last, char) {
			return false
		}

		if last == char {
			hasDouble = true
		}

		if isVowel(char) {
			vowels++
		}

		last = char
	}

	if vowels >= 3 && hasDouble {
		nice = true
	}

	return nice
}

func isVowel(char rune) bool {
	for _, v := range []rune{'a', 'e', 'i', 'o', 'u'} {
		if char == v {
			return true
		}
	}

	return false
}

func isBadString(last rune, current rune) bool {
	str := string(last) + string(current)
	for _, v := range []string{"ab", "cd", "pq", "xy"} {
		if str == v {
			return true
		}
	}

	return false
}
