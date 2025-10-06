package year2024day03

import (
	_ "embed"
	"regexp"
	"strconv"

	"github.com/polarfish/advent-of-code-go/puzzles/registry"
)

//go:embed year2024day03.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/3
	registry.AddPuzzle(2024, 3, "Mull It Over", input, part1, part2)
}

func part1(input string) string {
	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	matches := re.FindAllStringSubmatch(input, -1)
	var result int
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		result += left * right
	}
	return strconv.Itoa(result)
}

func part2(input string) string {
	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")
	matches := re.FindAllStringSubmatch(input, -1)
	var result int
	doMultiply := true
	for _, match := range matches {
		switch match[0][2] {
		case 'l': // mul
			if doMultiply {
				left, _ := strconv.Atoi(match[1])
				right, _ := strconv.Atoi(match[2])
				result += left * right
			}
		case 'n': // don't
			doMultiply = false
		default: // do
			doMultiply = true
		}
	}

	return strconv.Itoa(result)
}
