package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cpalsa/advent-of-code/go/util"
)

const inputFile = "input.txt"

type path struct {
	dest string
	dist int
}

func parsePaths(input []byte) map[string][]path {
	paths := make(map[string][]path)

	for _, p := range strings.Split(string(input), "\n") {
		pieces := strings.Split(p, " ")
		start := pieces[0]
		end := pieces[2]
		dist, err := strconv.Atoi(pieces[4])
		if err != nil {
			util.HandleFatal(err)
		}

		// make sure both location mappings exist and have their list of destinations initialized
		if _, ok := paths[start]; !ok {
			paths[start] = []path{}
		}
		if _, ok := paths[end]; !ok {
			paths[end] = []path{}
		}

		paths[start] = append(paths[start], path{end, dist})
		paths[end] = append(paths[end], path{start, dist})
	}

	return paths
}

func getShortestPath(paths []path) path {
	if len(paths) == 0 {
		util.HandleFatal(errors.New("no paths to compare"))
	}
	if len(paths) == 1 {
		return paths[0]
	}

	shortest := paths[0]
	for _, path := range paths {
		if path.dist < shortest.dist {
			shortest = path
		}
	}

	return shortest
}

func getLongestPath(paths []path) path {
	if len(paths) == 0 {
		util.HandleFatal(errors.New("no paths to compare"))
	}
	if len(paths) == 1 {
		return paths[0]
	}

	longest := paths[0]
	for _, path := range paths {
		if path.dist > longest.dist {
			longest = path
		}
	}

	return longest
}

func main() {
	input, err := util.LoadInput(inputFile)
	util.HandleFatal(err)

	paths := parsePaths(input)

	// Part 1
	shortestRoutes := make(map[string][]path)
	shortestDistances := make(map[string]int)
	for start, choices := range paths {
		shortestRoutes[start] = []path{}
		visited := []string{start}

		for len(choices) != 0 {
			// get the shortest path from our choices
			shortest := getShortestPath(choices)
			shortestRoutes[start] = append(shortestRoutes[start], shortest)
			shortestDistances[start] += shortest.dist
			visited = append(visited, shortest.dest)

			// filter available choices of routes by the names of places we already visited
			choices = util.FilterSlice(paths[shortest.dest], func(i int, v path) bool {
				return !util.SliceContains(visited, v.dest)
			})
		}
	}
	fmt.Printf("(Part 1) The shortest route Santa can take is a distance of %v\n", util.Min(util.MapVals(shortestDistances)...))

	longestRoutes := make(map[string][]path)
	longestDistances := make(map[string]int)
	for start, choices := range paths {
		longestRoutes[start] = []path{}
		visited := []string{start}

		for len(choices) != 0 {
			// get the longest path from our choices
			longest := getLongestPath(choices)
			longestRoutes[start] = append(longestRoutes[start], longest)
			longestDistances[start] += longest.dist
			visited = append(visited, longest.dest)

			// filter available choices of routes by the names of places we already visited
			choices = util.FilterSlice(paths[longest.dest], func(i int, v path) bool {
				return !util.SliceContains(visited, v.dest)
			})
		}
	}
	fmt.Printf("(Part 2) The longest route Santa can take is a distance of %v\n", util.Max(util.MapVals(longestDistances)...))
}
