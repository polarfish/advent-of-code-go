package year2024day02

import (
	_ "embed"
	"strings"

	"github.com/polarfish/advent-of-code-go/registry"
	"github.com/polarfish/advent-of-code-go/utils"
)

//go:embed year2024day02.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/2
	registry.AddSolution(2024, 2, "Red-Nosed Reports", input, part1, part2)
}

func part1(input string) string {
	result := 0
	lines := utils.Lines(input)
	for _, line := range lines {
		split := strings.Split(line, " ")
		levels := make([]int, len(split))
		for i, s := range split {
			levels[i] = utils.ToInt(s)
		}
		if findBadLevel(levels, -1) == -1 {
			result++
		}
	}

	return utils.ToStr(result)
}

func part2(input string) string {
	result := 0
	lines := utils.Lines(input)
	for _, line := range lines {
		split := strings.Split(line, " ")
		levels := make([]int, len(split))
		for i, s := range split {
			levels[i] = utils.ToInt(s)
		}
		badLevel := findBadLevel(levels, -1)
		if badLevel == -1 ||
			findBadLevel(levels, badLevel-2) == -1 ||
			findBadLevel(levels, badLevel-1) == -1 ||
			findBadLevel(levels, badLevel) == -1 {
			result++
		}
	}

	return utils.ToStr(result)
}

func findBadLevel(levels []int, levelToSkip int) int {
	var prev, delta int
	value := -1
	totalIncreasing := 0
	totalDecreasing := 0

	for i := 0; i < len(levels); i++ {
		if i == levelToSkip {
			continue
		}
		prev = value
		value = levels[i]
		if prev == -1 {
			continue
		}

		if value > prev {
			totalIncreasing++
			delta = value - prev
			if totalDecreasing > 0 {
				return i
			}
		} else if value < prev {
			totalDecreasing++
			delta = prev - value
			if totalIncreasing > 0 {
				return i
			}
		} else {
			return i
		}

		if delta > 3 {
			return i
		}
	}

	return -1
}
