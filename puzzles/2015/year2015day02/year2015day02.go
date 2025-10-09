package year2015day02

import (
	_ "embed"

	"github.com/polarfish/advent-of-code-go/registry"
	"github.com/polarfish/advent-of-code-go/utils"
)

//go:embed year2015day02.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/2
	registry.AddPuzzle(2015, 2, "I Was Told There Would Be No Math", input, part1, part2)
}

func part1(input string) string {
	dimensions := parseInput(input)
	var result int
	for _, d := range dimensions {
		l, w, h := d[0], d[1], d[2]
		s1 := l * w
		s2 := w * h
		s3 := h * l
		result += s1*2 + s2*2 + s3*2 + min(s1, s2, s3)
	}
	return utils.ToStr(result)
}

func part2(input string) string {
	dimensions := parseInput(input)
	var result int
	for _, d := range dimensions {
		l, w, h := d[0], d[1], d[2]
		result += (l+w+h-max(l, w, h))*2 + l*w*h
	}
	return utils.ToStr(result)
}

func parseInput(input string) [][]int {
	result := make([][]int, 0, 1000)
	ln := make([]int, 3)
	i := 0
	for _, ch := range input {
		if ch == 'x' {
			i++
		} else if ch == '\n' {
			result = append(result, ln)
			ln = make([]int, 3)
			i = 0
		} else {
			ln[i] = ln[i]*10 + int(ch) - 48
		}
	}
	return result
}
