package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

// Way to implement an enum in go
type Thing string

const (
	Snow Thing = "."
	Tree Thing = "#"
)

func parseLine(line string) []Thing {
	res := []Thing{}
	for i := 0; i < len(line); i++ {
		if line[i] == '#' {
			res = append(res, Tree)
		}
		if line[i] == '.' {
			res = append(res, Snow)
		}
	}
	return res
}

func parseInput(input string) [][]Thing {
	res := [][]Thing{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		res = append(res, parseLine(line))
	}
	return res
}

func calcTreeHits(input [][]Thing) int {
	res := 0
	colllength := len(input[0])
	for row, line := range input {
		if line[(row*3)%colllength] == Tree {
			res += 1
		}
	}
	return res
}

func main() {
	parsed := parseInput(getInput())
	fmt.Println("Hi there ", len(parsed), parsed[0])
	fmt.Println("Tree Hits: ", calcTreeHits(parsed))
}
