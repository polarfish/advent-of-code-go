package year2025day03

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day03.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/3
	registry.AddSolution(2025, 3, "Lobby", input, part1, part2)
}

func part1(input string) (string, error) {
	return strconv.FormatInt(calculateJoltage(input, 2), 10), nil
}

func part2(input string) (string, error) {
	return strconv.FormatInt(calculateJoltage(input, 12), 10), nil
}

func calculateJoltage(input string, batteriesCount int) int64 {
	var result int64 = 0
	lines := utils.Lines(input)
	for _, line := range lines {
		var bankJoltage int64 = 0
		startIndex := 0
		for j := 0; j < batteriesCount; j++ {
			var mx byte = 0
			for i := startIndex; i <= len(line)-batteriesCount+j; i++ {
				voltage := line[i] - '0'
				if voltage > mx {
					mx = voltage
					startIndex = i + 1
				}
			}
			bankJoltage = bankJoltage*10 + int64(mx)
		}
		result += bankJoltage
	}
	return result
}
