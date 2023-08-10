package main

import (
	"fmt"

	"github.com/cpalsa/advent-of-code/go/2015/util"
)

const inputFile = "input.txt"

const up = '^'
const down = 'v'
const left = '<'
const right = '>'

type position struct {
	x int
	y int
}

func (p position) move(direction rune) position {
	switch direction {
	case up:
		p.y += 1
	case down:
		p.y -= 1
	case left:
		p.x -= 1
	case right:
		p.x += 1
	}

	return p
}

func main() {
	input, err := util.LoadInput(inputFile)
	util.HandleFatal(err)

	// let's assume santa starts at 0,0
	pos := position{0, 0}

	// a map of all the house/positions santa delivers to and how many presents they got
	deliveries := make(map[position]int)

	// santa always delivers a present to his starting position
	deliveries[pos]++

	for _, d := range string(input) {
		// move to new position
		pos = pos.move(d)

		// deliver a present
		deliveries[pos]++
	}

	fmt.Printf("Working alone, Santa delivered presents to %v houses\n", len(deliveries))

	// Both human santa and robo-santa start at 0,0
	santaPos := position{0, 0}
	robotPos := position{0, 0}

	// a map of all the house/positions santa and robo-santa delivers to and how many presents they got
	deliveries = make(map[position]int)

	// santa and robo-santa always delivers a present to their starting position
	deliveries[santaPos]++
	deliveries[robotPos]++

	santaMoves := true
	for _, d := range string(input) {
		moverPos := &santaPos

		if !santaMoves {
			moverPos = &robotPos
		}

		// move to new position
		*moverPos = moverPos.move(d)

		// deliver a present
		deliveries[*moverPos]++

		// toggle between human santa and robo-santa
		santaMoves = !santaMoves
	}

	fmt.Printf("Working together, Santa and Robo-Santa delivered presents to %v houses\n", len(deliveries))
}
