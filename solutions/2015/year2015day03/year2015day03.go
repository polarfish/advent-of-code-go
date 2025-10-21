package year2015day03

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2015day03.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/3
	registry.AddSolution(2015, 3, "Perfectly Spherical Houses in a Vacuum", input, part1, part2)
}

func part1(input string) (string, error) {
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

	return strconv.Itoa(len(m)), nil
}

func part2(input string) (string, error) {
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

	return strconv.Itoa(len(m)), nil
}
