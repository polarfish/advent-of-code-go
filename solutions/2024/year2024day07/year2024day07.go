package year2024day07

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day07.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/7
	registry.AddSolution(2024, 7, "Bridge Repair", input, part1, part2)
}

func part1(input string) (string, error) {
	var result int64
	for _, line := range utils.Lines(input) {
		split := strings.Split(line, ":")
		expected, _ := strconv.ParseInt(split[0], 10, 64)
		valuesStr := strings.Fields(split[1])
		values := make([]int, len(valuesStr))
		for i, v := range valuesStr {
			values[i], _ = strconv.Atoi(v)
		}
		if solve1(expected, int64(values[0]), values, 1) {
			result += expected
		}
	}
	return strconv.FormatInt(result, 10), nil
}

func part2(input string) (string, error) {
	var result int64
	for _, line := range utils.Lines(input) {
		split := strings.Split(line, ":")
		expected, _ := strconv.ParseInt(split[0], 10, 64)
		valuesStr := strings.Fields(split[1])
		values := make([]int, len(valuesStr))
		for i, v := range valuesStr {
			values[i], _ = strconv.Atoi(v)
		}
		if solve2(expected, int64(values[0]), values, 1) {
			result += expected
		}
	}
	return strconv.FormatInt(result, 10), nil
}

func solve1(expected, current int64, values []int, i int) bool {
	if i == len(values) {
		return expected == current
	}
	if current > expected {
		return false
	}
	return solve1(expected, current+int64(values[i]), values, i+1) ||
		solve1(expected, current*int64(values[i]), values, i+1)
}

func solve2(expected, current int64, values []int, i int) bool {
	if i == len(values) {
		return expected == current
	}
	if current > expected {
		return false
	}
	return solve2(expected, current+int64(values[i]), values, i+1) ||
		solve2(expected, current*int64(values[i]), values, i+1) ||
		solve2(expected, concat(current, values[i]), values, i+1)
}

func concat(a int64, b int) int64 {
	if b == 0 {
		a *= 10
	} else {
		tempB := b
		for tempB > 0 {
			tempB /= 10
			a *= 10
		}
	}
	return a + int64(b)
}
