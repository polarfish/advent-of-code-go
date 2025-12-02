package year2025day02

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestYear2025Day02Part1(t *testing.T) {
	test.Assert(t, "1227775554", part1, sample)
	test.Assert(t, "31210613313", part1, input)
}

func TestYear2025Day02Part2(t *testing.T) {
	test.Assert(t, "4174379265", part2, sample)
	test.Assert(t, "41823587546", part2, input)
}
