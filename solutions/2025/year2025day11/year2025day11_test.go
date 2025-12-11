package year2025day11

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`

var sample2 = `
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`

func TestYear2025Day11Part1Sample(t *testing.T) {
	test.Assert(t, "5", part1, sample)
}

func TestYear2025Day11Part1(t *testing.T) {
	test.Assert(t, "566", part1, input)
}

func TestYear2025Day11Part2Sample(t *testing.T) {
	test.Assert(t, "2", part2, sample2)
}

func TestYear2025Day11Part2(t *testing.T) {
	test.Assert(t, "331837854931968", part2, input)
}
