package year2025day05

import (
	"cmp"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day05.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/5
	registry.AddSolution(2025, 5, "Cafeteria", input, part1, part2)
}

func part1(input string) (string, error) {
	result := 0

	lines := utils.Lines(input)
	ranges, ids := parseInput(lines, true)
	mergedRanges := mergeRanges(ranges)

	lookUpSlice := make([]int64, 0, len(mergedRanges)*2)
	for _, rng := range mergedRanges {
		lookUpSlice = append(lookUpSlice, rng[0], rng[1])
	}

	for _, id := range ids {
		idx, ok := slices.BinarySearch(lookUpSlice, id)
		if ok || idx&1 == 1 {
			result++
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	result := int64(0)
	lines := utils.Lines(input)
	ranges, _ := parseInput(lines, true)
	mergedRanges := mergeRanges(ranges)

	for _, rng := range mergedRanges {
		result = result + rng[1] - rng[0] + 1
	}
	return strconv.FormatInt(result, 10), nil
}

func parseInput(lines []string, withIds bool) ([][]int64, []int64) {
	var ranges = make([][]int64, 0)
	var ids = make([]int64, 0)
	parsingRanges := true
	for _, line := range lines {
		if line == "" {
			if withIds {
				parsingRanges = false
				continue
			}
			break
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			rng := make([]int64, 2)
			rng[0], _ = strconv.ParseInt(parts[0], 10, 64)
			rng[1], _ = strconv.ParseInt(parts[1], 10, 64)
			ranges = append(ranges, rng)
		} else {
			id, _ := strconv.ParseInt(line, 10, 64)
			ids = append(ids, id)
		}
	}
	return ranges, ids
}

func mergeRanges(ranges [][]int64) [][]int64 {
	slices.SortFunc(ranges, func(s1, s2 []int64) int {
		return cmp.Compare(s1[0], s2[0])
	})
	var mergedRanges = make([][]int64, 0, len(ranges))
	var last = []int64{ranges[0][0], ranges[0][1]}
	mergedRanges = append(mergedRanges, last)
	for i := 1; i < len(ranges); i++ {
		curr := ranges[i]
		if curr[0] <= last[1] {
			last[1] = max(curr[1], last[1])
		} else {
			last = []int64{curr[0], curr[1]}
			mergedRanges = append(mergedRanges, last)
		}
	}
	return mergedRanges
}
