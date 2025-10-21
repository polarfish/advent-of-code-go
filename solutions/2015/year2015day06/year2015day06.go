package year2015day06

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2015day06.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/6
	registry.AddSolution(2015, 6, "Probably a Fire Hazard", input, part1, part2)
}

func part1(input string) (string, error) {
	return simulateLights(input,
		func(i int8) int8 { return 1 },
		func(i int8) int8 { return 0 },
		func(i int8) int8 { return i ^ 1 })
}

func part2(input string) (string, error) {
	return simulateLights(input,
		func(i int8) int8 { return i + 1 },
		func(i int8) int8 { return max(0, i-1) },
		func(i int8) int8 { return i + 2 })
}

func simulateLights(input string, turnOn func(i int8) int8, turnOff func(i int8) int8, toggle func(i int8) int8) (string, error) {
	var err error
	const size = 1000
	matrix := make([][]int8, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int8, size)
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		var from, to []string
		var x1, y1, x2, y2 int
		var op func(int8) int8
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

		x1, err = strconv.Atoi(from[0])
		if err != nil {
			return "", err
		}
		y1, err = strconv.Atoi(from[1])
		if err != nil {
			return "", err
		}
		x2, err = strconv.Atoi(to[0])
		if err != nil {
			return "", err
		}
		y2, err = strconv.Atoi(to[1])
		if err != nil {
			return "", err
		}

		for j := y1; j <= y2; j++ {
			for i := x1; i <= x2; i++ {
				matrix[j][i] = op(matrix[j][i])
			}
		}
	}

	result := 0
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			result += int(matrix[j][i])
		}
	}

	return strconv.Itoa(result), nil
}
