package year2024day11

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2024day11.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/11
	registry.AddSolution(2024, 11, "Plutonian Pebbles", input, part1, part2)
}

func part1(input string) (string, error) {
	stones := parseInput(input)
	result := calculateStones(stones, 25)
	return strconv.FormatInt(result, 10), nil
}

func part2(input string) (string, error) {
	stones := parseInput(input)
	result := calculateStones(stones, 75)
	return strconv.FormatInt(result, 10), nil
}

func parseInput(input string) []int64 {
	fields := strings.Fields(strings.TrimSpace(input))
	stones := make([]int64, len(fields))
	for i, f := range fields {
		v, _ := strconv.ParseInt(f, 10, 64)
		stones[i] = v
	}
	return stones
}

func calculateStones(stones []int64, steps int) int64 {
	frequency := make(map[int64]int64)
	for _, s := range stones {
		frequency[s]++
	}
	for i := 0; i < steps; i++ {
		newFrequency := make(map[int64]int64, len(frequency)*2)
		for num, count := range frequency {
			if num == 0 {
				newFrequency[1] += count
			} else if splitRes := split(num); splitRes != nil {
				newFrequency[splitRes[0]] += count
				newFrequency[splitRes[1]] += count
			} else {
				newFrequency[num*2024] += count
			}
		}
		frequency = newFrequency
	}
	var sum int64
	for _, v := range frequency {
		sum += v
	}
	return sum
}

func split(l int64) []int64 {
	buf := l
	digits := 0
	for buf > 0 {
		digits++
		buf /= 10
	}
	if digits%2 == 0 {
		multiplier := pow(10, int64(digits/2), 1)
		return []int64{l / multiplier, l % multiplier}
	}
	return nil
}

func pow(n, p, accumulator int64) int64 {
	if p == 0 {
		return accumulator
	}
	return pow(n, p-1, accumulator*n)
}
