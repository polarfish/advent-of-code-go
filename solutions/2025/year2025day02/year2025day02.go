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
	return solve(input, checkInvalidPart1)
}

func part2(input string) (string, error) {
	return solve(input, checkInvalidPart2)
}

func solve(input string, checkInvalid func(int) bool) (string, error) {
	result := 0
	ranges := strings.Split(strings.TrimSpace(input), ",")
	for _, rng := range ranges {
		var err error
		var from, to int
		pair := strings.Split(rng, "-")
		from, err = strconv.Atoi(pair[0])
		if err != nil {
			return "", utils.ErrBadInput
		}
		to, err = strconv.Atoi(pair[1])
		if err != nil {
			return "", utils.ErrBadInput
		}
		for i := from; i <= to; i++ {
			match := checkInvalid(i)

			if match {
				result += i
			}
		}
	}
	return strconv.Itoa(result), nil
}

func checkInvalidPart1(n int) bool {
	id := strconv.Itoa(n)
	l := len(id)

	if l%2 != 0 {
		return false
	}

	size := l / 2

	for j := 0; j < size; j++ {
		if id[j] != id[len(id)-size+j] {
			return false
		}
	}

	return true
}

func checkInvalidPart2(n int) bool {
	id := strconv.Itoa(n)
	l := len(id)
	maxSize := l / 2

top:
	for size := maxSize; size > 0; size-- {
		if l%size != 0 {
			continue
		}
		repeats := l / size
		for j := 0; j < size; j++ {
			target := id[j]
			for k := 1; k < repeats; k++ {
				if id[j+k*size] != target {
					continue top
				}
			}
		}
		return true
	}

	return false
}
