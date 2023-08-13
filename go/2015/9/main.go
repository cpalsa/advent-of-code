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
	// A graph may be a better "long term" solution if we were to need to expand
	paths := make(map[string][]path)

	for _, p := range strings.Split(string(input), "\n") {
		pieces := strings.Split(p, " ")
		start := pieces[0]
		end := pieces[2]
		dist := util.ExitIfError(strconv.Atoi(pieces[4]))

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

func getShortestPath(paths []path) (path, error) {
	if len(paths) == 0 {
		return path{}, errors.New("no paths to compare")
	}
	if len(paths) == 1 {
		return paths[0], nil
	}

	shortest := paths[0]
	for _, path := range paths {
		if path.dist < shortest.dist {
			shortest = path
		}
	}

	return shortest, nil
}

func getLongestPath(paths []path) (path, error) {
	if len(paths) == 0 {
		return path{}, errors.New("no paths to compare")
	}
	if len(paths) == 1 {
		return paths[0], nil
	}

	longest := paths[0]
	for _, path := range paths {
		if path.dist > longest.dist {
			longest = path
		}
	}

	return longest, nil
}

func main() {
	input := util.ExitIfError(util.LoadInput(inputFile))

	paths := parsePaths(input)

	// Part 1
	shortestRoutes := make(map[string][]path)
	shortestDistances := make(map[string]int)
	for start, choices := range paths {
		shortestRoutes[start] = []path{}
		visited := []string{start}

		for len(choices) != 0 {
			// get the shortest path from our choices
			shortest := util.ExitIfError(getShortestPath(choices))
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
			longest := util.ExitIfError(getLongestPath(choices))
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
