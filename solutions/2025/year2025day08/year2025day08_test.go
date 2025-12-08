package year2025day08

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

func TestYear2025Day08Part1Sample(t *testing.T) {
	test.Assert(t, "40", part1, sample)
}

func TestYear2025Day08Part1(t *testing.T) {
	test.Assert(t, "54600", part1, input)
}

func TestYear2025Day08Part2Sample(t *testing.T) {
	test.Assert(t, "25272", part2, sample)
}

func TestYear2025Day08Part2(t *testing.T) {
	test.Assert(t, "107256172", part2, input)
}
