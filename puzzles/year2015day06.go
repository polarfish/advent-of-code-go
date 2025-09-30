package puzzles

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed year2015day06.txt
var year2015Day06Input string

func init() {
	// https://adventofcode.com/2015/day/6
	addPuzzle(2015, 6, "Probably a Fire Hazard", year2015Day06Input, year2015Day06Part1, year2015Day06Part2)
}

func year2015Day06Part1(input string) string {
	return simulateLights(input,
		func(i int) int { return 1 },
		func(i int) int { return 0 },
		func(i int) int { return i ^ 1 })
}

func year2015Day06Part2(input string) string {
	return simulateLights(input,
		func(i int) int { return i + 1 },
		func(i int) int { return max(0, i-1) },
		func(i int) int { return i + 2 })
}

func simulateLights(input string, turnOn func(i int) int, turnOff func(i int) int, toggle func(i int) int) string {
	const size = 1000
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		var from, to []string
		var x1, y1, x2, y2 int
		var op func(int) int
		if split[1] == "off" {
			op = turnOff
			from = strings.Split(split[2], ",")
			to = strings.Split(split[4], ",")
		} else if split[1] == "on" {
			op = turnOn
			from = strings.Split(split[2], ",")
			to = strings.Split(split[4], ",")
		} else { // toggle
			op = toggle
			from = strings.Split(split[1], ",")
			to = strings.Split(split[3], ",")
		}

		x1, _ = strconv.Atoi(from[0])
		y1, _ = strconv.Atoi(from[1])
		x2, _ = strconv.Atoi(to[0])
		y2, _ = strconv.Atoi(to[1])

		for j := y1; j <= y2; j++ {
			for i := x1; i <= x2; i++ {
				matrix[j][i] = op(matrix[j][i])
			}
		}
	}

	result := 0
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			result += matrix[j][i]
		}
	}

	return strconv.Itoa(result)
}
