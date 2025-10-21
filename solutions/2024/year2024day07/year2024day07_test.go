package year2024day07

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day07Part1(t *testing.T) {
	test.Assert(t, "21572148763543", part1, input)
}

func TestYear2024Day07Part2(t *testing.T) {
	test.Assert(t, "581941094529163", part2, input)
}
