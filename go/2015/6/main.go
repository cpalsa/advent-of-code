package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

type position struct {
	x int
	y int
}

type command struct {
	action string
	start  position
	end    position
}

func main() {
	input, err := util.LoadInput(inputFile)
	util.HandleFatal(err)

	commands := parseCommands(input)

	// only records lights that are ON
	lightMap := make(map[position]bool)
	// only records lights with brightness > 0
	brightnessMap := make(map[position]int)

	for _, cmd := range commands {
		for y := cmd.start.y; y <= cmd.end.y; y++ {
			for x := cmd.start.x; x <= cmd.end.x; x++ {
				pos := position{x, y}

				// if the map key doesn't exist, the light is off (false)
				status := lightMap[pos]

				// if the map key doesn't exist, brightness is 0
				brightness := brightnessMap[pos]

				switch cmd.action {
				case "toggle":
					status = !status
					brightness += 2
				case "on":
					status = true
					brightness += 1
				case "off":
					status = false
					brightness -= 1
				}

				// if light is being turned off remove it from map
				if !status {
					delete(lightMap, pos)
				} else {
					lightMap[pos] = status
				}

				if brightness <= 0 {
					delete(brightnessMap, pos)
				} else {
					brightnessMap[pos] = brightness
				}
			}
		}
	}
	fmt.Printf("(Part 1) There are %v lights turned on\n", len(lightMap))

	brightness := 0
	for _, b := range brightnessMap {
		brightness += b
	}

	fmt.Printf("(Part 2) The total brightness is %v\n", brightness)
}

func parseCommands(input []byte) []command {
	cmds := []command{}

	for _, row := range strings.Split(string(input), "\n") {
		cmd := command{}
		str := strings.Split(row, " ")

		if strings.HasPrefix(row, "toggle") {
			cmd.action = "toggle"
			cmd.start = parsePosition(str[1])
			cmd.end = parsePosition(str[3])
			cmds = append(cmds, cmd)
			continue
		}

		if strings.HasPrefix(row, "turn off") {
			cmd.action = "off"
		} else {
			cmd.action = "on"
		}

		cmd.start = parsePosition(str[2])
		cmd.end = parsePosition(str[4])
		cmds = append(cmds, cmd)
	}

	return cmds
}

func parsePosition(str string) position {
	var err error
	pos := position{}
	coords := strings.Split(string(str), ",")

	pos.x, err = strconv.Atoi(coords[0])
	if err != nil {
		util.HandleFatal(err)
	}

	pos.y, err = strconv.Atoi(coords[1])
	if err != nil {
		util.HandleFatal(err)
	}

	return pos
}
