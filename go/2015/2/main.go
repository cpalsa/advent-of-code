package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

func parseDimensions(input []byte) ([][]int, error) {
	var dimensions [][]int

	// split by each new line of the input and iterate
	for _, line := range strings.Split(string(input), "\n") {
		var dims []int

		// split each line of LxWxH into a slice of ints
		for _, val := range strings.Split(line, "x") {
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			dims = append(dims, num)
		}

		dimensions = append(dimensions, dims)
	}

	return dimensions, nil
}

func main() {
	input := util.ExitIfError(util.LoadInput(inputFile))

	// [][l, w, h]
	dimensions := util.ExitIfError(parseDimensions(input))

	paper := 0
	ribbon := 0
	for _, dims := range dimensions {
		minArea := min(dims[0]*dims[1], dims[0]*dims[2], dims[1]*dims[2])
		paper += (2 * dims[0] * dims[1]) + (2 * dims[1] * dims[2]) + (2 * dims[2] * dims[0]) + minArea

		// sort dims ascending, we don't care about keep position anymore for shortest perimeter
		sort.Ints(dims)
		ribbon += (2 * dims[0]) + (2 * dims[1]) + (dims[0] * dims[1] * dims[2])
	}

	fmt.Printf("(Part 1) The elves require %v sqft of wrapping paper\n", paper)
	fmt.Printf("(Part 2) The elves require %v feet of ribbon\n", ribbon)
}
