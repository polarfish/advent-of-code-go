package year2015day02

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/utils"
)

//go:embed day02.txt
var input string

func New() *utils.Puzzle {
	return &utils.Puzzle{
		Year:  2015,
		Day:   2,
		Name:  "I Was Told There Would Be No Math",
		Input: &input,
		Part1: Part1,
		Part2: Part2,
	}
}

func Part1(input *string) string {
	dimensions := parseInput(input)
	var result int
	for _, d := range dimensions {
		l, w, h := d[0], d[1], d[2]
		s1 := l * w
		s2 := w * h
		s3 := h * l
		result += s1*2 + s2*2 + s3*2 + min(s1, s2, s3)
	}
	return strconv.Itoa(result)
}

func Part2(input *string) string {
	dimensions := parseInput(input)
	var result int
	for _, d := range dimensions {
		l, w, h := d[0], d[1], d[2]
		result += (l+w+h-max(l, w, h))*2 + l*w*h
	}
	return strconv.Itoa(result)
}

func parseInput(input *string) [][]int {
	result := make([][]int, 0, strings.Count(*input, "\n"))
	ln := make([]int, 3)
	i := 0
	for _, ch := range *input {
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
