package year2015day03

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/utils"
)

func TestYear2015day03Part1(t *testing.T) {
	utils.Test(t, "2572", part1(input))
}

func TestYear2015day03Part2(t *testing.T) {
	utils.Test(t, "2631", part2(input))
}
