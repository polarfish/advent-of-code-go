package year2024day04

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day04.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/4
	registry.AddSolution(2024, 4, "Ceres Search", input, part1, part2)
}

func part1(input string) (string, error) {
	m := buildMap(input)

	result := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			result += tryMatch(m, x, y)
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	m := buildMap(input)

	result := 0
	for y := 1; y < len(m)-1; y++ {
		for x := 1; x < len(m[y])-1; x++ {
			if m[y][x] == 'A' &&
				((m[y-1][x-1] == 'M' && m[y+1][x+1] == 'S') ||
					(m[y-1][x-1] == 'S' && m[y+1][x+1] == 'M')) &&
				((m[y-1][x+1] == 'M' && m[y+1][x-1] == 'S') ||
					(m[y-1][x+1] == 'S' && m[y+1][x-1] == 'M')) {
				result++
			}
		}
	}

	return strconv.Itoa(result), nil
}

const Xmas = "XMAS"

func buildMap(input string) [][]byte {
	lines := utils.Lines(input)
	m := make([][]byte, len(lines))
	for y, line := range lines {
		m[y] = []byte(line)
	}
	return m
}

func tryMatch(m [][]byte, x, y int) int {
	var total int
	for dir := 0; dir < 8; dir++ {
		total += tryMatchDirection(m, x, y, dir)
	}
	return total
}

func tryMatchDirection(m [][]byte, x, y, direction int) int {
	dy, dx := 0, 0

	switch direction {
	case 0, 1, 7:
		dy = -1
	case 3, 4, 5:
		dy = 1
	}
	switch direction {
	case 1, 2, 3:
		dx = 1
	case 5, 6, 7:
		dx = -1
	}

	currentY, currentX := y, x
	for _, ch := range []byte(Xmas) {
		if currentX < 0 || currentX >= len(m[0]) || currentY < 0 || currentY >= len(m) {
			return 0
		}

		if m[currentY][currentX] != ch {
			return 0
		}

		currentY += dy
		currentX += dx
	}

	return 1
}
