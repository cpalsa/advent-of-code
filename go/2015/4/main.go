package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

func main() {
	input := util.ExitIfError(util.LoadInput(inputFile))

	num := 0
	hash := ""
	for !strings.HasPrefix(hash, "00000") {
		num++
		test := string(input) + fmt.Sprint(num)
		hash = fmt.Sprintf("%x", md5.Sum([]byte(test)))
	}

	fmt.Printf("(Part 1) The lowest positive integer that yields a hash starting with five zeroes (%v) is %v\n", hash, num)

	num = 0
	hash = ""
	for !strings.HasPrefix(hash, "000000") {
		num++
		test := string(input) + fmt.Sprint(num)
		hash = fmt.Sprintf("%x", md5.Sum([]byte(test)))
	}

	fmt.Printf("(Part 2) The lowest positive integer that yields a hash starting with six zeroes (%v) is %v\n", hash, num)
}
