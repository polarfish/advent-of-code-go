package year2024day18

import (
	"container/list"
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day18.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/18
	registry.AddSolution(2024, 18, "RAM Run", input, part1, part2)
}

const gridSize = 71

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type pos struct {
	x, y, steps int
}

type fallingByte struct {
	x, y int
}

func parseInput(input string) ([]fallingByte, error) {
	lines := utils.Lines(input)
	res := make([]fallingByte, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			return nil, utils.ErrBadInput
		}
		x, err1 := strconv.Atoi(split[0])
		y, err2 := strconv.Atoi(split[1])
		if err1 != nil || err2 != nil {
			return nil, utils.ErrBadInput
		}
		res[i] = fallingByte{x, y}
	}
	return res, nil
}

func part1(input string) (string, error) {
	bytes, err := parseInput(input)
	if err != nil {
		return "", err
	}
	var mapGrid = utils.NewGrid[int](gridSize, gridSize)
	for i := 0; i < 1024 && i < len(bytes); i++ {
		b := bytes[i]
		mapGrid[b.y][b.x] = 1
	}
	steps := bfs(mapGrid)
	return strconv.Itoa(steps), nil
}

func bfs(mapGrid [][]int) int {
	var visited = utils.NewGrid[int](gridSize, gridSize)
	queue := list.New()
	queue.PushBack(pos{0, 0, 0})
	visited[0][0] = 1
	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(pos)
		if p.x == gridSize-1 && p.y == gridSize-1 {
			return p.steps
		}
		for _, d := range dirs {
			nx, ny := p.x+d[0], p.y+d[1]
			if nx >= 0 && nx < gridSize && ny >= 0 && ny < gridSize && mapGrid[ny][nx] == 0 && visited[ny][nx] == 0 {
				visited[ny][nx] = 1
				queue.PushBack(pos{nx, ny, p.steps + 1})
			}
		}
	}
	return 0
}

func part2(input string) (string, error) {
	bytes, err := parseInput(input)
	if err != nil {
		return "", err
	}
	var mapGrid = utils.NewGrid[int](gridSize, gridSize)
	var visited = utils.NewGrid[int](gridSize, gridSize)
	utils.ResetGrid(visited, 2)
	var blocking fallingByte
	for i, b := range bytes {
		mapGrid[b.y][b.x] = 1
		if i >= 1024 {
			if visited[b.y][b.x] == 2 {
				utils.ResetGrid(visited, 0)
				visited[0][0] = 1
				res := dfs(mapGrid, visited, 0, 0, 0) // with good directions dfs is faster than bfs here
				if res == 0 {
					blocking = b
					break
				}
			}
		}
	}
	if blocking.x == 0 && blocking.y == 0 {
		return "", nil
	}
	return strconv.Itoa(blocking.x) + "," + strconv.Itoa(blocking.y), nil
}

func dfs(mapGrid [][]int, visited [][]int, x, y, steps int) int {
	if x == gridSize-1 && y == gridSize-1 {
		visited[y][x] = 2
		return steps
	}
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && nx < gridSize && ny >= 0 && ny < gridSize && mapGrid[ny][nx] == 0 && visited[ny][nx] == 0 {
			visited[ny][nx] = 1
			res := dfs(mapGrid, visited, nx, ny, steps+1)
			if res != 0 {
				visited[y][x] = 2
				return res
			}
		}
	}
	return 0
}
