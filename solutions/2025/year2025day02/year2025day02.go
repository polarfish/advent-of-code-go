package year2025day02

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day02.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/2
	registry.AddSolution(2025, 2, "Gift Shop", input, part1, part2)
}

func part1(input string) (string, error) {
	return solve(input,
		createInvalidChecker([][]int{{}, {}, {1}, {}, {2}, {}, {3}, {}, {4}, {}, {5}}))
}

func part2(input string) (string, error) {
	return solve(input,
		createInvalidChecker([][]int{{}, {}, {1}, {1}, {2, 1}, {1}, {3, 2, 1}, {1}, {4, 2, 1}, {3, 1}, {5, 2, 1}}))
}

func solve(input string, checkInvalid func(int, int) bool) (string, error) {
	result := 0
	ranges := strings.Split(strings.TrimSpace(input), ",")
	for i := 0; i < len(ranges); i++ {
		rng := ranges[i]
		var err error
		var from, to int
		pair := strings.Split(rng, "-")

		numSize := len(pair[0])
		if len(pair[0]) < len(pair[1]) {
			ranges = append(ranges, "1"+strings.Repeat("0", numSize)+"-"+pair[1])
			pair[1] = strings.Repeat("9", len(pair[0]))
		}

		from, err = strconv.Atoi(pair[0])
		if err != nil {
			return "", utils.ErrBadInput
		}
		to, err = strconv.Atoi(pair[1])
		if err != nil {
			return "", utils.ErrBadInput
		}

		for i := from; i <= to; i++ {
			if checkInvalid(numSize, i) {
				result += i
			}
		}
	}
	return strconv.Itoa(result), nil
}

func createInvalidChecker(matrixSizes [][]int) func(int, int) bool {
	return func(numSize, n int) bool {
	top:
		for _, size := range matrixSizes[numSize] {
			div := 1
			for range size {
				div *= 10
			}
			t := n
			target := t % div
			t /= div
			for t > 0 {
				if t%div != target {
					continue top
				}
				t = t / div
			}
			return true
		}

		return false
	}

}
