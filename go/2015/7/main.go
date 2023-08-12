package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

type circuit map[string]uint16
type command struct {
	wire1 string
	wire2 string
	gate  string
	dest  string
}

func parseCommands(input []byte) []command {
	cmds := []command{}

	for _, c := range strings.Split(string(input), "\n") {
		cmd := command{}
		pieces := strings.Split(c, " ")

		if len(pieces) == 3 {
			cmd.wire1 = pieces[0]
		}
		if len(pieces) == 4 {
			cmd.gate = pieces[0]
			cmd.wire1 = pieces[1]
		}
		if len(pieces) == 5 {
			cmd.wire1 = pieces[0]
			cmd.gate = pieces[1]
			cmd.wire2 = pieces[2]
		}

		cmd.dest = pieces[len(pieces)-1]

		cmds = append(cmds, cmd)
	}

	return cmds
}

func (cmd command) eval(c circuit) (ok bool) {
	// first check circuit if a signal already exists at the output destination
	_, ok = c[cmd.dest]
	if ok {
		return ok
	}

	resolveSignal := func(wire string) (uint16, bool) {
		parsed, err := strconv.ParseUint(wire, 10, 16)
		if err == nil {
			// no error = raw uint16 signal
			return uint16(parsed), true
		}

		// read signal from circuit
		sig, ok := c[wire]

		return sig, ok
	}

	switch cmd.gate {
	case "AND":
		sig1, ok1 := resolveSignal(cmd.wire1)
		sig2, ok2 := resolveSignal(cmd.wire2)

		if ok1 && ok2 {
			ok = true
			c[cmd.dest] = sig1 & sig2
		}
	case "OR":
		sig1, ok1 := resolveSignal(cmd.wire1)
		sig2, ok2 := resolveSignal(cmd.wire2)

		if ok1 && ok2 {
			ok = true
			c[cmd.dest] = sig1 | sig2
		}
	case "LSHIFT":
		sig1, ok1 := resolveSignal(cmd.wire1)
		sig2, ok2 := resolveSignal(cmd.wire2)

		if ok1 && ok2 {
			ok = true
			c[cmd.dest] = sig1 << sig2
		}
	case "RSHIFT":
		sig1, ok1 := resolveSignal(cmd.wire1)
		sig2, ok2 := resolveSignal(cmd.wire2)

		if ok1 && ok2 {
			ok = true
			c[cmd.dest] = sig1 >> sig2
		}
	case "NOT":
		sig1, ok1 := resolveSignal(cmd.wire1)

		if ok1 {
			ok = true
			c[cmd.dest] = ^sig1
		}
	default:
		sig1, ok1 := resolveSignal(cmd.wire1)

		if ok1 {
			ok = true
			c[cmd.dest] = sig1
		}
	}

	return ok
}

func (c circuit) eval(commands []command) {
	for len(commands) != 0 {
		cmd := commands[0]
		ok := cmd.eval(c)

		// one of the input wires doesn't have a signal to resolve yet
		if !ok {
			// try again later
			commands = append(commands, cmd)
		}

		// advance to the next command
		commands = commands[1:]
	}
}

func main() {
	input, err := util.LoadInput(inputFile)
	util.HandleFatal(err)

	commands := parseCommands(input)
	circuit := make(circuit)
	circuit.eval(commands)
	fmt.Printf("(Part 1) The signal %v is provided to wire 'a'\n", circuit["a"])

	clear(circuit)
	circuit["b"] = 46065
	circuit.eval(commands)
	fmt.Printf("(Part 2) The new signal %v is provided to wire 'a'\n", circuit["a"])
}
