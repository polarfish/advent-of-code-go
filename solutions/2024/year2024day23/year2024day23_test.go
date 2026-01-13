package year2024day23

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day23Part1(t *testing.T) {
	test.Assert(t, "1163", part1, input)
}

func TestYear2024Day23Part2(t *testing.T) {
	test.Assert(t, "bm,bo,ee,fo,gt,hv,jv,kd,md,mu,nm,wx,xh", part2, input)
}
