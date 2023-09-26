package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func parseLine(lineString string) Line {
	partStrings := strings.Split(lineString, " -> ")
	x1, err := strconv.Atoi(strings.Split(partStrings[0], ",")[0])
	if err != nil {
		log.Fatal("This should never ever happen.")
	}
	y1, err := strconv.Atoi(strings.Split(partStrings[0], ",")[1])
	if err != nil {
		log.Fatal("This should never ever happen.")
	}
	x2, err := strconv.Atoi(strings.Split(partStrings[1], ",")[0])
	if err != nil {
		log.Fatal("This should never ever happen.")
	}
	y2, err := strconv.Atoi(strings.Split(partStrings[1], ",")[1])
	if err != nil {
		log.Fatal("This should never ever happen.")
	}

	return Line{Point{x1, y1}, Point{x2, y2}}
}

func shouldInclude(pointOne Point, pointTwo Point) bool {
	return pointOne.x == pointTwo.x || pointOne.y == pointTwo.y
}

func main() {
	lines := strings.Split(getInput(), "\n")
	for _, lineStr := range lines {
		line := parseLine(lineStr)
		if shouldInclude(line.p1, line.p2) {
			fmt.Println(line)
		}
	}
}
