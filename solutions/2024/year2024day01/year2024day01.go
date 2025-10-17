package year2024day01

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day01.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/1
	registry.AddSolution(2024, 1, "Historian Hysteria", input, part1, part2)
}

func part1(input string) string {
	lines := utils.Lines(input)
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		split := strings.Split(line, "   ")
		left[i], right[i] = utils.ToInt(split[0]), utils.ToInt(split[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	var total int
	for i := 0; i < len(lines); i++ {
		total += utils.Abs(left[i] - right[i])
	}

	return strconv.Itoa(total)
}

func part2(input string) string {
	lines := utils.Lines(input)
	left := make([]int, len(lines))
	right := make(map[int]int, len(lines))
	for i, line := range lines {
		split := strings.Split(line, "   ")
		l, r := utils.ToInt(split[0]), utils.ToInt(split[1])
		left[i] = l
		right[r]++
	}

	var total int
	for _, n := range left {
		if val, exists := right[n]; exists {
			total += n * val
		}
	}

	return strconv.Itoa(total)
}
