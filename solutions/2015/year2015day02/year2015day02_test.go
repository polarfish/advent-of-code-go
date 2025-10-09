package year2015day02

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/utils"
)

func TestYear2015day02Part1(t *testing.T) {
	utils.Test(t, "1606483", part1(input))
}

func TestYear2015day02Part2(t *testing.T) {
	utils.Test(t, "3842356", part2(input))
}
