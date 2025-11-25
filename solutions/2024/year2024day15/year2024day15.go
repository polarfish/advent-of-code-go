package year2024day15

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day15.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/15
	registry.AddSolution(2024, 15, "Warehouse Woes", input, part1, part2)
}

func part1(input string) (string, error) {
	split := strings.SplitN(input, "\n\n", 2)
	if len(split) < 2 {
		return "", utils.ErrBadInput
	}
	maps := parseMap(split[0])
	moves := []rune(strings.ReplaceAll(split[1], "\n", ""))
	robotX, robotY := locateRobot(maps)

	for _, dir := range moves {
		x, y := robotX, robotY
		step := 0
		for {
			step++
			x = nextX(x, dir)
			y = nextY(y, dir)
			if maps[y][x] == '.' {
				maps[robotY][robotX] = '.'
				if step > 1 {
					maps[y][x] = 'O'
					robotY = nextY(robotY, dir)
					robotX = nextX(robotX, dir)
					maps[robotY][robotX] = '@'
				} else {
					robotY = y
					robotX = x
					maps[robotY][robotX] = '@'
				}
				break
			} else if maps[y][x] == '#' {
				break
			}
		}
	}
	result := sumUpCoordinates(maps, 'O')
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	split := strings.SplitN(input, "\n\n", 2)
	if len(split) < 2 {
		return "", utils.ErrBadInput
	}
	maps := parseWideMap(split[0])
	moves := []rune(strings.ReplaceAll(split[1], "\n", ""))
	robotX, robotY := locateRobot(maps)

	for _, dir := range moves {
		if canMove(maps, robotX, robotY, dir) {
			move(maps, robotX, robotY, dir)
			robotX = nextX(robotX, dir)
			robotY = nextY(robotY, dir)
		}
	}
	result := sumUpCoordinates(maps, '[')
	return strconv.Itoa(result), nil
}

func sumUpCoordinates(maps [][]rune, symbol rune) int {
	result := 0
	for y := range maps {
		for x := range maps[0] {
			if maps[y][x] == symbol {
				result += y*100 + x
			}
		}
	}
	return result
}

func parseMap(mapString string) [][]rune {
	lines := utils.Lines(mapString)
	m := make([][]rune, len(lines))
	for i, line := range lines {
		m[i] = []rune(line)
	}
	return m
}

func parseWideMap(mapString string) [][]rune {
	maps := parseMap(mapString)
	wideMap := make([][]rune, len(maps))
	for y := range maps {
		wideMap[y] = make([]rune, len(maps[0])*2)
		for x := range maps[0] {
			switch maps[y][x] {
			case '#':
				wideMap[y][x*2] = '#'
				wideMap[y][x*2+1] = '#'
			case '.':
				wideMap[y][x*2] = '.'
				wideMap[y][x*2+1] = '.'
			case 'O':
				wideMap[y][x*2] = '['
				wideMap[y][x*2+1] = ']'
			case '@':
				wideMap[y][x*2] = '@'
				wideMap[y][x*2+1] = '.'
			default:
				wideMap[y][x*2] = '.'
				wideMap[y][x*2+1] = '.'
			}
		}
	}
	return wideMap
}

func locateRobot(maps [][]rune) (int, int) {
	for y := range maps {
		for x := range maps[0] {
			if maps[y][x] == '@' {
				return x, y
			}
		}
	}
	return -1, -1
}

func nextX(x int, dir rune) int {
	if dir == '>' {
		return x + 1
	} else if dir == '<' {
		return x - 1
	}
	return x
}

func nextY(y int, dir rune) int {
	if dir == 'v' {
		return y + 1
	} else if dir == '^' {
		return y - 1
	}
	return y
}

func canMove(maps [][]rune, x, y int, dir rune) bool {
	x2 := nextX(x, dir)
	y2 := nextY(y, dir)
	if maps[y2][x2] == '.' {
		return true
	}
	if dir == '^' || dir == 'v' {
		if maps[y2][x2] == '[' {
			return canMove(maps, x2, y2, dir) && canMove(maps, x2+1, y2, dir)
		} else if maps[y2][x2] == ']' {
			return canMove(maps, x2-1, y2, dir) && canMove(maps, x2, y2, dir)
		}
	} else if dir == '>' && maps[y2][x2] == '[' {
		return canMove(maps, x2+1, y2, dir)
	} else if dir == '<' && maps[y2][x2] == ']' {
		return canMove(maps, x2-1, y2, dir)
	}
	return false
}

func move(maps [][]rune, x, y int, dir rune) {
	x2 := nextX(x, dir)
	y2 := nextY(y, dir)
	if dir == '^' || dir == 'v' {
		if maps[y2][x2] == '[' {
			move(maps, x2, y2, dir)
			move(maps, x2+1, y2, dir)
		} else if maps[y2][x2] == ']' {
			move(maps, x2-1, y2, dir)
			move(maps, x2, y2, dir)
		}
	} else if dir == '>' && maps[y2][x2] == '[' {
		move(maps, x2+1, y2, dir)
		maps[y2][x2+1] = maps[y2][x2]
	} else if dir == '<' && maps[y2][x2] == ']' {
		move(maps, x2-1, y2, dir)
		maps[y2][x2-1] = maps[y2][x2]
	}
	maps[y2][x2] = maps[y][x]
	maps[y][x] = '.'
}
