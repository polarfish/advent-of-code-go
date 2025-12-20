package year2024day20

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day20.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/20
	registry.AddSolution(2024, 20, "Race Condition", input, part1, part2)
}

const (
	X = 0
	Y = 1
)

type Step struct {
	x, y int
	time int
}

func part1(input string) (string, error) {
	lines := utils.Lines(input)
	h := len(lines)
	w := len(lines[0])

	mapGrid := utils.NewGrid[byte](h, w)

	startX, startY := -1, -1
	endX, endY := -1, -1

	for y := range h {
		for x := range w {
			switch lines[y][x] {
			case 'S':
				startX, startY = x, y
				mapGrid[y][x] = '.'
			case 'E':
				endX, endY = x, y
				mapGrid[y][x] = '.'
			default:
				mapGrid[y][x] = lines[y][x]
			}
		}
	}

	path := make([]*Step, 0)
	visited := utils.NewGrid[*Step](h, w)

	curr := &Step{startX, startY, 0}
	path = append(path, curr)
	visited[startY][startX] = curr

	for !(curr.x == endX && curr.y == endY) {
		x, y := curr.x, curr.y

		// up
		if y > 0 && mapGrid[y-1][x] != '#' && visited[y-1][x] == nil {
			next := &Step{x, y - 1, curr.time + 1}
			curr = next
			path = append(path, curr)
			visited[y-1][x] = path[len(path)-1]
			continue
		}

		// right
		if x < w-1 && mapGrid[y][x+1] != '#' && visited[y][x+1] == nil {
			next := &Step{x + 1, y, curr.time + 1}
			curr = next
			path = append(path, curr)
			visited[y][x+1] = path[len(path)-1]
			continue
		}

		// down
		if y < h-1 && mapGrid[y+1][x] != '#' && visited[y+1][x] == nil {
			next := &Step{x, y + 1, curr.time + 1}
			curr = next
			path = append(path, curr)
			visited[y+1][x] = path[len(path)-1]
			continue
		}

		// left
		if x > 0 && mapGrid[y][x-1] != '#' && visited[y][x-1] == nil {
			next := &Step{x - 1, y, curr.time + 1}
			curr = next
			path = append(path, curr)
			visited[y][x-1] = path[len(path)-1]
			continue
		}

		return "", utils.ErrBadInput
	}

	shortcuts := make(map[int]int)

	for _, step := range path {
		x, y := step.x, step.y

		// up
		if y > 1 && mapGrid[y-1][x] == '#' {
			if hs := visited[y-2][x]; hs != nil {
				if saved := hs.time - step.time - 2; saved > 0 {
					shortcuts[saved]++
				}
			}
		}

		// right
		if x < w-2 && mapGrid[y][x+1] == '#' {
			if hs := visited[y][x+2]; hs != nil {
				if saved := hs.time - step.time - 2; saved > 0 {
					shortcuts[saved]++
				}
			}
		}

		// down
		if y < h-2 && mapGrid[y+1][x] == '#' {
			if hs := visited[y+2][x]; hs != nil {
				if saved := hs.time - step.time - 2; saved > 0 {
					shortcuts[saved]++
				}
			}
		}

		// left
		if x > 1 && mapGrid[y][x-1] == '#' {
			if hs := visited[y][x-2]; hs != nil {
				if saved := hs.time - step.time - 2; saved > 0 {
					shortcuts[saved]++
				}
			}
		}
	}

	result := 0
	for saved, count := range shortcuts {
		if saved >= 100 {
			result += count
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	mapGrid := parseMap(input)
	result, err := calculateShortcuts(mapGrid, 20)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(result), nil
}

func parseMap(input string) [][]byte {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func calculateShortcuts(mapGrid [][]byte, shortcutMaxSize int) (int, error) {
	h := len(mapGrid)
	w := len(mapGrid[0])

	start := [2]int{-1, -1}
	end := [2]int{-1, -1}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			switch mapGrid[y][x] {
			case 'S':
				start[X], start[Y] = x, y
				mapGrid[y][x] = '.'
			case 'E':
				end[X], end[Y] = x, y
				mapGrid[y][x] = '.'
			}
		}
	}

	path := make([][2]int, 0)
	curr := start
	prev := [2]int{-1, -1}
	path = append(path, curr)

	for !(curr[X] == end[X] && curr[Y] == end[Y]) {
		x, y := curr[X], curr[Y]

		// up
		if y > 0 && mapGrid[y-1][x] == '.' && !(x == prev[X] && y-1 == prev[Y]) {
			prev = curr
			curr = [2]int{x, y - 1}
			path = append(path, curr)
			continue
		}

		// right
		if x < w-1 && mapGrid[y][x+1] != '#' && !(x+1 == prev[X] && y == prev[Y]) {
			prev = curr
			curr = [2]int{x + 1, y}
			path = append(path, curr)
			continue
		}

		// down
		if y < h-1 && mapGrid[y+1][x] != '#' && !(x == prev[X] && y+1 == prev[Y]) {
			prev = curr
			curr = [2]int{x, y + 1}
			path = append(path, curr)
			continue
		}

		// left
		if x > 0 && mapGrid[y][x-1] != '#' && !(x-1 == prev[X] && y == prev[Y]) {
			prev = curr
			curr = [2]int{x - 1, y}
			path = append(path, curr)
			continue
		}

		return 0, utils.ErrBadInput
	}

	result := 0
	for i := 0; i < len(path)-1; i++ {
		for j := i + 1; j < len(path); j++ {
			from := path[i]
			to := path[j]

			manhattan := int(math.Abs(float64(from[X]-to[X])) +
				math.Abs(float64(from[Y]-to[Y])))
			trackDist := j - i

			if manhattan <= shortcutMaxSize && trackDist-manhattan >= 100 {
				result++
			}
		}
	}

	return result, nil
}
