package year2025day07

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day07.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/7
	registry.AddSolution(2025, 7, "Laboratories", input, part1, part2)
}

func part1(input string) (string, error) {
	result := 0
	lines := utils.Lines(input)
	beams := make([]int, len(lines[0]))

	for i, ch := range lines[0] {
		if ch == 'S' {
			beams[i] = 1
		}
	}

	for l := 1; l < len(lines); l++ {
		for i, ch := range lines[l] {
			if ch == '^' && beams[i] == 1 {
				result += 1
				beams[i] = 0
				if i > 0 {
					beams[i-1] = 1
				}
				if i < len(lines[0])-1 {
					beams[i+1] = 1
				}
			}
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	var result int64 = 0
	lines := utils.Lines(input)
	data := make([][]byte, len(lines))
	for i := range data {
		data[i] = []byte(lines[i])
	}

	timeline := make([]int64, len(data[0]))
	for y, line := range data {
		for x, ch := range line {
			switch ch {
			case 'S':
				data[y+1][x] = '|'
				timeline[x] = 1
			case '^':
				if y-1 >= 0 && data[y-1][x] == '|' {
					data[y][x-1] = '|'
					data[y][x+1] = '|'
					timeline[x-1] += timeline[x]
					timeline[x+1] += timeline[x]
					timeline[x] = 0
				}
			case '.':
				if y-1 >= 0 && data[y-1][x] == '|' {
					data[y][x] = '|'
				}
			}
		}
	}

	for _, val := range timeline {
		result += val
	}

	return strconv.FormatInt(result, 10), nil
}
