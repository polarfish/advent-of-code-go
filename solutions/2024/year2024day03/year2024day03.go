package year2024day03

import (
	_ "embed"
	"regexp"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day03.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/3
	registry.AddSolution(2024, 3, "Mull It Over", input, part1, part2)
}

func part1(input string) string {
	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	matches := re.FindAllStringSubmatch(input, -1)
	var result int
	for _, match := range matches {
		result += utils.ToInt(match[1]) * utils.ToInt(match[2])
	}
	return utils.ToStr(result)
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
				result += utils.ToInt(match[1]) * utils.ToInt(match[2])
			}
		case 'n': // don't
			doMultiply = false
		default: // do
			doMultiply = true
		}
	}

	return utils.ToStr(result)
}
