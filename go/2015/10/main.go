package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/cpalsa/advent-of-code/go/util"
	"github.com/dlclark/regexp2"
)

const inputFile = "input.txt"

// perform the look-and-say sequence for n iterations
// WAAAAAAY too slow -- regex version below
// func lookAndSay(input string, n uint) string {
// 	if n == 0 {
// 		return input
// 	}

// 	look := []rune(input)
// 	say := ""

// 	buf := []rune{}
// 	for len(look) != 0 {
// 		char := look[0]

// 		// start of new sequence or continuation of current sequence
// 		if len(buf) == 0 || char == buf[0] {
// 			buf = append(buf, char)
// 		}

// 		// end of string or next char ends the sequence
// 		if len(look[1:]) == 0 || look[1:][0] != char {
// 			// process the buffer
// 			say += fmt.Sprintf("%d%v", len(buf), string(char))
// 			buf = []rune{}
// 		}

// 		// advance to next char
// 		look = look[1:]
// 	}

// 	return lookAndSay(say, n-1)
// }

// perform the look-and-say sequence for n iterations
func lookAndSay(input string, n uint) string {
	if n == 0 {
		return input
	}

	// built-in regex engine doesnt support backreferences
	replaceFn := func(m regexp2.Match) string {
		return fmt.Sprintf("%d%v", len(m.Runes()), string(m.Runes()[0]))
	}
	regex := regexp2.MustCompile(`(\d)\1*`, 0) //regexp2.RegexOptions
	result := util.ExitIfError(regex.ReplaceFunc(input, replaceFn, -1, -1))

	return lookAndSay(result, n-1)
}

func main() {
	input := util.ExitIfError(util.LoadInput(inputFile))

	las40 := lookAndSay((string(input)), 40)
	ln40 := utf8.RuneCountInString(las40)
	fmt.Printf("(Part 1) the length of the look-and-say of '%v' at 40 iterations is: %v\n", string(input), ln40)

	ln50 := utf8.RuneCountInString(lookAndSay(las40, 10))
	fmt.Printf("(Part 2) the length of the look-and-say of '%v' at 50 iterations is: %v\n", string(input), ln50)
}
