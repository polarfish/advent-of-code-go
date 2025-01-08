package year2015day02

import (
	_ "embed"
	"regexp"
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

var dimensionsRegexp = regexp.MustCompile("(\\d+)x(\\d+)x(\\d+)")

func parseInput(input *string) [][]int {
	linesCount := strings.Count(*input, "\n") + 1
	result := make([][]int, 0, linesCount)
	matches := dimensionsRegexp.FindAllStringSubmatch(*input, -1)
	for _, m := range matches {
		l, _ := strconv.Atoi(m[1])
		w, _ := strconv.Atoi(m[2])
		h, _ := strconv.Atoi(m[3])
		result = append(result, []int{l, w, h})
	}
	return result
}
