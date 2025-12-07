package year2025day07

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day07.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/7
	registry.AddSolution(2025, 7, "Laboratories", input, part1, part2)
}

func part1(input string) (string, error) {
	result := 0
	lines := utils.Lines(input)
	beams := make([]int, len(lines[0]))

	for i, ch := range lines[0] {
		if ch == 'S' {
			beams[i] = 1
		}
	}

	for l := 1; l < len(lines); l++ {
		for i, ch := range lines[l] {
			if ch == '^' && beams[i] == 1 {
				result += 1
				beams[i] = 0
				if i > 0 {
					beams[i-1] = 1
				}
				if i < len(lines[0])-1 {
					beams[i+1] = 1
				}
			}
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	var result int64 = 0
	lines := utils.Lines(input)
	var beam int

	for i, ch := range lines[0] {
		if ch == 'S' {
			beam = i
		}
	}

	memo := make([][]int64, len(lines))
	for i := range memo {
		memo[i] = make([]int64, len(lines[0]))
	}

	result = countTimelines(lines, 0, beam, memo)

	return strconv.FormatInt(result, 10), nil
}

func countTimelines(lines []string, i int, beam int, memo [][]int64) int64 {
	if beam < 0 || beam >= len(lines[0]) {
		return 0
	}

	if i >= len(lines) {
		return 1
	}

	if memo[i][beam] > 0 {
		return memo[i][beam]
	}

	var timelines int64

	if lines[i][beam] == '^' {
		timelines = countTimelines(lines, i, beam-1, memo) + countTimelines(lines, i, beam+1, memo)
	} else {
		timelines = countTimelines(lines, i+1, beam, memo)
	}

	memo[i][beam] = timelines
	return timelines
}
