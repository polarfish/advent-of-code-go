package year2025day06

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func TestYear2025Day06Part1Sample(t *testing.T) {
	test.Assert(t, "4277556", part1, sample)
}

func TestYear2025Day06Part1(t *testing.T) {
	test.Assert(t, "5667835681547", part1, input)
}

func TestYear2025Day06Part2Sample(t *testing.T) {
	test.Assert(t, "3263827", part2, sample)
}

func TestYear2025Day06Part2(t *testing.T) {
	test.Assert(t, "9434900032651", part2, input)
}
