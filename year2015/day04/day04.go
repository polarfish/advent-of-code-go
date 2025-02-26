package year2015day04

import (
	"crypto/md5"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/utils"
)

//go:embed day04.txt
var input string

const maxIterations = 100_000_000

func New() *utils.Puzzle {
	return &utils.Puzzle{
		Year:  2015,
		Day:   4,
		Name:  "The Ideal Stocking Stuffer",
		Input: input,
		Part1: Part1,
		Part2: Part2,
	}
}

func Part1(input string) string {
	return solve(input, func(result [16]byte) bool {
		return result[0] == 0 && result[1] == 0 && result[2] < 16
	})
}

func Part2(input string) string {
	return solve(input, func(result [16]byte) bool {
		return result[0] == 0 && result[1] == 0 && result[2] == 0
	})
}

func solve(input string, predicate func([16]byte) bool) string {
	inputBytes := []byte(strings.TrimSpace(input))

	var result [16]byte
	for i := 1; i < maxIterations; i++ {
		result = md5.Sum(slices.Concat(inputBytes, []byte(strconv.Itoa(i))))
		if predicate(result) {
			return strconv.Itoa(i)
		}
	}
	return utils.ERR
}
