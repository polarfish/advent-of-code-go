package year2024day06

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day06.txt
var input string

func init() {
	registry.AddSolution(2024, 6, "Guard Gallivant", input, part1, part2)
}

var turn = [4]int{1, 2, 3, 0}

func part1(input string) (string, error) {
	grid := parseMap(input)
	guardX, guardY := findGuardPosition(grid)
	visits := make([][]int, len(grid))
	for i := range visits {
		visits[i] = make([]int, len(grid[0]))
	}
	checkForLoop(grid, visits, guardX, guardY, 0, false)
	result := 0
	for _, row := range visits {
		for _, v := range row {
			if v != 0 {
				result++
			}
		}
	}
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	grid := parseMap(input)
	guardX, guardY := findGuardPosition(grid)
	visits := make([][]int, len(grid))
	for i := range visits {
		visits[i] = make([]int, len(grid[0]))
	}
	result := checkForLoop(grid, visits, guardX, guardY, 0, true)
	return strconv.Itoa(result), nil
}

func parseMap(input string) [][]byte {
	lines := utils.Lines(input)
	m := make([][]byte, len(lines))
	for i, line := range lines {
		m[i] = []byte(line)
	}
	return m
}

func findGuardPosition(grid [][]byte) (int, int) {
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == '^' {
				return x, y
			}
		}
	}
	panic("Guard not found")
}

func copyVisits(visits [][]int, visitsCopy [][]int) [][]int {
	for i := range visits {
		copy(visitsCopy[i], visits[i])
	}
	return visitsCopy
}

func checkForLoop(grid [][]byte, visits [][]int, x, y, dir int, tryObstruction bool) int {
	var x2, y2, result int
	var visitsCopy [][]int
	if tryObstruction {
		visitsCopy = make([][]int, len(visits))
		for i := range visits {
			visitsCopy[i] = make([]int, len(visits[0]))
		}
	}
	for {
		visits[y][x] |= 1 << dir
		switch dir {
		case 0: // up
			x2, y2 = x, y-1
			if y2 < 0 {
				return result
			}
		case 1: // right
			x2, y2 = x+1, y
			if x2 >= len(grid[0]) {
				return result
			}
		case 2: // down
			x2, y2 = x, y+1
			if y2 >= len(grid) {
				return result
			}
		case 3: // left
			x2, y2 = x-1, y
			if x2 < 0 {
				return result
			}
		default:
			panic("Unknown direction")
		}

		if grid[y2][x2] == '#' {
			dir = turn[dir]
		} else {
			if (visits[y2][x2] & (1 << dir)) > 0 {
				return result + 1
			}
			if tryObstruction && visits[y2][x2] == 0 {
				grid[y2][x2] = '#'
				result += checkForLoop(grid, copyVisits(visits, visitsCopy), x, y, turn[dir], false)
				grid[y2][x2] = '.'
			}
			x, y = x2, y2
		}
	}
}
