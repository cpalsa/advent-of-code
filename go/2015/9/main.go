package main

import (
	"fmt"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

func main() {
	input, err := util.LoadInput(inputFile)
	util.HandleFatal(err)

	fmt.Printf("%v\n", string(input))
}
