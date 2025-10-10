package year2015day07

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/utils"
)

func TestYear2015Day07Part1(t *testing.T) {
	utils.Test(t, "3176", part1(input))
}

func TestYear2015Day07Part2(t *testing.T) {
	utils.Test(t, "14710", part2(input))
}
