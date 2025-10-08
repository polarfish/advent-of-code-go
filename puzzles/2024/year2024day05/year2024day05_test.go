package year2024day05

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/puzzles/utils"
)

func TestYear2024Day05Part1(t *testing.T) {
	utils.Test(t, "5713", part1(input))
}

func TestYear2024Day05Part2(t *testing.T) {
	utils.Test(t, "5180", part2(input))
}
