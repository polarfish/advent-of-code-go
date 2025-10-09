package year2015day04

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/utils"
)

func TestYear2015day04Part1(t *testing.T) {
	utils.Test(t, "282749", part1(input))
}

func TestYear2015day04Part2(t *testing.T) {
	utils.Test(t, "9962624", part2(input))
}
