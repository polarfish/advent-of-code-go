package year2025day01

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestYear2025Day01Part1(t *testing.T) {
	test.Assert(t, "3", part1, sample)
	test.Assert(t, "1040", part1, input)
}

func TestYear2025Day01Part2(t *testing.T) {
	test.Assert(t, "6", part2, sample)
	test.Assert(t, "6027", part2, input)
}
