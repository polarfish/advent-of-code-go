package puzzles

import (
	"crypto/md5"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed year2015day04.txt
var year2015Day04Input string

const maxIterations = 100_000_000

func init() {
	// https://adventofcode.com/2015/day/4
	addPuzzle(2015, 4, "The Ideal Stocking Stuffer", year2015Day04Input, year2015Day04Part1, year2015Day04Part2)
}

func year2015Day04Part1(input string) string {
	return solve(input, func(result [16]byte) bool {
		return result[0] == 0 && result[1] == 0 && result[2] < 16
	})
}

func year2015Day04Part2(input string) string {
	return solve(input, func(result [16]byte) bool {
		return result[0] == 0 && result[1] == 0 && result[2] == 0
	})
}

func solve(input string, test func([16]byte) bool) string {
	inputBytes := append([]byte(strings.TrimSpace(input)), 0)
	threshold := 10
	var l = len(inputBytes) - 1
	var j, k int
	var res [16]byte
	for i := 1; i < maxIterations; i++ {
		if i >= threshold {
			l++
			threshold *= 10
			inputBytes = append(inputBytes, 0)
		}
		j, k = i, l
		for j > 0 {
			inputBytes[k] = byte(j%10) + 48
			k -= 1
			j /= 10
		}

		res = md5.Sum(inputBytes)
		if test(res) {
			return strconv.Itoa(i)
		}
	}
	return errorResult
}
