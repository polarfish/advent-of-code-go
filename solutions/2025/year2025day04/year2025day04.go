package year2025day04

import (
	"container/list"
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day04.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/4
	registry.AddSolution(2025, 4, "Printing Department", input, part1, part2)
}

var neighborOffsets = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func part1(input string) (string, error) {
	result := 0
	lines := utils.Lines(input)

	h, w := len(lines), len(lines[0])
	m := make([][]int, h)
	for y := 0; y < h; y++ {
		m[y] = make([]int, w)
		for x := 0; x < w; x++ {
			if lines[y][x] == '@' {
				m[y][x] = 1
			}
		}
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if m[y][x] == 0 {
				continue
			}

			adjacent := 0
			for _, off := range neighborOffsets {
				adjacent += safeGet(m, y+off[0], x+off[1], h, w)
			}

			if adjacent < 4 {
				result++
			}
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	result := 0
	lines := utils.Lines(input)

	h, w := len(lines), len(lines[0])
	m := make([][]int, h)
	heat := make([][]int, h)
	for y := 0; y < h; y++ {
		m[y] = make([]int, w)
		heat[y] = make([]int, w)
		for x := 0; x < w; x++ {
			if lines[y][x] == '@' {
				m[y][x] = 1
			}
		}
	}

	removalQueue := list.New()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if m[y][x] == 0 {
				continue
			}
			adjacent := 0
			for _, off := range neighborOffsets {
				adjacent += safeGet(m, y+off[0], x+off[1], h, w)
			}
			heat[y][x] = adjacent
			if adjacent < 4 {
				removalQueue.PushBack([]int{y, x})
			}
		}
	}

	for removalQueue.Len() > 0 {
		coords := removalQueue.Remove(removalQueue.Front()).([]int)
		y, x := coords[0], coords[1]
		m[y][x] = 0
		result++

		for _, off := range neighborOffsets {
			reduceAdjacency(m, heat, y+off[0], x+off[1], h, w, removalQueue)
		}
	}

	return strconv.Itoa(result), nil
}

func safeGet(m [][]int, y, x, h, w int) int {
	if x >= 0 && x < w && y >= 0 && y < h {
		return m[y][x]
	}
	return 0
}

func reduceAdjacency(m, heat [][]int, y, x, h, w int, removalQueue *list.List) {
	if x >= 0 && x < w && y >= 0 && y < h && m[y][x] == 1 {
		heat[y][x]--
		if heat[y][x] == 3 {
			removalQueue.PushBack([]int{y, x})
		}
	}
}
