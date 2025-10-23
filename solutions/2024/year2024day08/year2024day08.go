package year2024day08

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day08.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/8
	registry.AddSolution(2024, 8, "Resonant Collinearity", input, part1, part2)
}

func part1(input string) (string, error) {
	w, h, groups := locateAntennas(input)
	antinodes := placeAntinodes(w, h, groups, false)
	return strconv.Itoa(len(antinodes)), nil
}

func part2(input string) (string, error) {
	w, h, groups := locateAntennas(input)
	antinodes := placeAntinodes(w, h, groups, true)
	return strconv.Itoa(len(antinodes)), nil
}

type antenna struct {
	y         int
	x         int
	frequency rune
}

type antinode struct {
	x int
	y int
}

func locateAntennas(input string) (int, int, map[rune][]antenna) {
	lines := utils.Lines(input)
	antennas := make([]antenna, 0)

	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			antennas = append(antennas, antenna{x, y, char})
		}
	}

	antennasGroups := make(map[rune][]antenna)
	for _, antenna := range antennas {
		antennasGroups[antenna.frequency] = append(antennasGroups[antenna.frequency], antenna)
	}

	return len(lines[0]), len(lines), antennasGroups
}

func placeAntinodes(width, height int, groups map[rune][]antenna, fullLine bool) map[antinode]struct{} {
	var x, y int

	antinodes := make(map[antinode]struct{})

	for _, group := range groups {
		for i := 0; i < len(group)-1; i++ {
			for j := i + 1; j < len(group); j++ {
				a1 := group[i]
				a2 := group[j]

				deltaY := abs(a1.x - a2.x)
				deltaX := abs(a1.y - a2.y)

				if a1.y > a2.y {
					a1, a2 = a2, a1
				}

				left := 1
				if fullLine {
					left = 0
				}

				for {
					left++

					newY := a2.x - deltaY*left
					if a2.x < a1.x {
						newY = a2.x + deltaY*left
					}

					x, y = a2.y-deltaX*left, newY
					if isOffBounds(x, y, height, width) {
						break
					}

					antinodes[antinode{x, y}] = struct{}{}

					if !fullLine {
						break
					}
				}

				right := 1
				if fullLine {
					right = 0
				}

				for {
					right++

					newY := a1.x + deltaY*right
					if a2.x < a1.x {
						newY = a1.x - deltaY*right
					}

					x = a1.y + deltaX*right
					y = newY

					if isOffBounds(x, y, height, width) {
						break
					}

					antinodes[antinode{x, y}] = struct{}{}

					if !fullLine {
						break
					}
				}
			}
		}
	}
	return antinodes
}

func isOffBounds(x, y, height, width int) bool {
	return x >= width || x < 0 || y >= height || y < 0
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
