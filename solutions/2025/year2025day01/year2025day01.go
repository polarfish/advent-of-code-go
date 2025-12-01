package year2025day01

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day01.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/1
	registry.AddSolution(2025, 1, "Secret Entrance", input, part1, part2)
}

func part1(input string) (string, error) {
	lines := utils.Lines(input)
	dial := 50
	result := 0
	for _, line := range lines {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", utils.ErrBadInput
		}

		if line[0] == 'R' {
			dial += num
		} else {
			dial -= num
		}

		for dial < 0 {
			dial += 100
		}

		dial %= 100

		if dial == 0 {
			result += 1
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	lines := utils.Lines(input)
	dial := 50
	result := 0
	for _, line := range lines {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", utils.ErrBadInput
		}

		result += num / 100
		num %= 100

		if line[0] == 'R' {
			dial += num
			if dial > 99 {
				dial -= 100
				result++
			}
		} else {
			oldDial := dial
			dial -= num
			if dial < 0 {
				dial += 100
				if oldDial != 0 {
					result++
				}

			}
			if dial == 0 {
				result++
			}
		}
	}

	return strconv.Itoa(result), nil
}
