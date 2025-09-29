package puzzles

import (
	_ "embed"
	"strconv"
)

//go:embed year2015day03.txt
var year2015Day03Input string

func init() {
	// https://adventofcode.com/2015/day/3
	addPuzzle(2015, 3, "Perfectly Spherical Houses in a Vacuum", year2015Day03Input, year2015Day03Part1, year2015Day03Part2)
}

func year2015Day03Part1(input string) string {
	x, y := 1<<15, 1<<15
	m := make(map[int]int, len(input)/2)
	m[(y<<16)|x] = 1
	for _, ch := range input {
		switch ch {
		case '^':
			y--
		case '>':
			x++
		case 'v':
			y++
		case '<':
			x--
		}
		m[(y<<16)|x] = 1
	}

	return strconv.Itoa(len(m))
}

func year2015Day03Part2(input string) string {
	coords := [4]int{1 << 15, 1 << 15, 1 << 15, 1 << 15} // 0, 1 - Santa; 2, 3 - Robo-Santa
	m := make(map[int]int, len(input)/2)
	m[(coords[0]<<16)|coords[1]] = 1
	for i, ch := range input {
		offset := (i & 1) << 1
		switch ch {
		case '^':
			coords[offset+1]--
		case '>':
			coords[offset]++
		case 'v':
			coords[offset+1]++
		case '<':
			coords[offset]--
		}
		m[(coords[offset]<<16)|coords[offset+1]] = 1
	}

	return strconv.Itoa(len(m))
}
