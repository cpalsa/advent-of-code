package main

import (
	"fmt"

	util "github.com/cpalsa/advent-of-code/go/2015/util"
)

const inputFile = "input.txt"

const up = '('
const down = ')'

func main() {
	input, err := util.LoadInput(inputFile)
	util.HandleFatal(err)

	floor, pos := 0, -1
	for i, direction := range string(input) {
		if direction == up {
			floor++
		}

		if direction == down {
			floor--
		}

		if floor == -1 && pos == -1 {
			pos = i + 1
		}
	}

	fmt.Printf("Santa has to go to floor %v\n", floor)
	fmt.Printf("Santa first enters the basement at position %v\n", pos)
}
