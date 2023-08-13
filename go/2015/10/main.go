package main

import (
	"fmt"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

func main() {
	input := util.ExitIfError(util.LoadInput(inputFile))

	fmt.Printf("%v", string(input))
}
