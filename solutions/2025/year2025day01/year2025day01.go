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

const startPosition = 50
const totalPositions = 100

func part1(input string) (string, error) {
	lines := utils.Lines(input)
	dial := startPosition
	result := 0
	for _, line := range lines {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", utils.ErrBadInput
		}

		switch line[0] {
		case 'R':
			dial += num
		case 'L':
			dial -= num
		default:
			return "", utils.ErrBadInput
		}

		for dial < 0 {
			dial += totalPositions
		}

		dial %= totalPositions

		if dial == 0 {
			result++
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	lines := utils.Lines(input)
	dial := startPosition
	result := 0
	for _, line := range lines {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", utils.ErrBadInput
		}

		result += num / totalPositions
		num %= totalPositions

		switch line[0] {
		case 'R':
			dial += num
			if dial >= totalPositions {
				dial -= totalPositions
				result++
			}
		case 'L':
			oldDial := dial
			dial -= num
			if dial < 0 {
				dial += totalPositions
				if oldDial != 0 {
					result++
				}
			}
			if dial == 0 {
				result++
			}
		default:
			return "", utils.ErrBadInput
		}
	}

	return strconv.Itoa(result), nil
}
