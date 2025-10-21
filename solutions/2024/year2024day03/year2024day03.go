package year2024day03

import (
	_ "embed"
	"regexp"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2024day03.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/3
	registry.AddSolution(2024, 3, "Mull It Over", input, part1, part2)
}

func part1(input string) (string, error) {
	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	matches := re.FindAllStringSubmatch(input, -1)
	var result int
	for _, match := range matches {
		value1, err := strconv.Atoi(match[1])
		if err != nil {
			return "", err
		}
		value2, err := strconv.Atoi(match[2])
		if err != nil {
			return "", err
		}
		result += value1 * value2
	}
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")
	matches := re.FindAllStringSubmatch(input, -1)
	var result int
	doMultiply := true
	for _, match := range matches {
		switch match[0][2] {
		case 'l': // mul
			if doMultiply {
				value1, err := strconv.Atoi(match[1])
				if err != nil {
					return "", err
				}
				value2, err := strconv.Atoi(match[2])
				if err != nil {
					return "", err
				}
				result += value1 * value2
			}
		case 'n': // don't
			doMultiply = false
		default: // do
			doMultiply = true
		}
	}

	return strconv.Itoa(result), nil
}
