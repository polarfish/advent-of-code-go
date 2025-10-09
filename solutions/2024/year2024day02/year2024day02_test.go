package year2024day02

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/utils"
)

func TestYear2024Day02Part1(t *testing.T) {
	utils.Test(t, "334", part1(input))
}

func TestYear2024Day02Part2(t *testing.T) {
	utils.Test(t, "400", part2(input))
}
