package year2015day06

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/puzzles/utils"
)

func TestYear2015Day06Part1(t *testing.T) {
	utils.Test(t, "400410", part1(input))
}

func TestYear2015Day06Part2(t *testing.T) {
	utils.Test(t, "15343601", part2(input))
}
