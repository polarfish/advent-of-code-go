package year2015day05

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2015day05.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/5
	registry.AddSolution(2015, 5, "Doesn't He Have Intern-Elves For This?", input, part1, part2)
}

func part1(input string) (string, error) {
	var result int
	lines := strings.Split(input, "\n")
top:
	for _, line := range lines {
		var vowels = 0
		var twice = false
		var prev int32
		for i, ch := range line {
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
				vowels++
			}
			if i > 0 {
				if ch == prev {
					twice = true
				}
				if (prev == 'a' && ch == 'b') || (prev == 'c' && ch == 'd') || (prev == 'p' && ch == 'q') || (prev == 'x' && ch == 'y') {
					continue top
				}
			}
			prev = ch
		}
		if vowels >= 3 && twice {
			result++
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	var result int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pairMap := make(map[string]int, len(line))
		var trio = false
		var pair = false
		var prev, prev2 int32
		for i, ch := range line {
			if i > 0 {
				key := string([]int32{prev, ch})
				val, ok := pairMap[key]
				if ok {
					if i-1-val > 1 {
						pair = true
					}
				} else {
					pairMap[key] = i - 1
				}
			}
			if i > 1 {
				if ch == prev2 {
					trio = true
				}
			}
			prev2 = prev
			prev = ch
		}
		if trio && pair {
			result++
		}
	}

	return strconv.Itoa(result), nil
}
