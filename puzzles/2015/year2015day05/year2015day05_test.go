package year2015day05

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/puzzles/utils"
)

func TestYear2015day05Part1(t *testing.T) {
	utils.Test(t, "255", part1(input))
}

func TestYear2015day05Part2(t *testing.T) {
	utils.Test(t, "55", part2(input))
}
