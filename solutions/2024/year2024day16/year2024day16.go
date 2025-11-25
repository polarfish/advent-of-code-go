package year2024day16

import (
	"container/heap"
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day16.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/16
	registry.AddSolution(2024, 16, "Reindeer Maze", input, part1, part2)
}

func part1(input string) (string, error) {
	return strconv.Itoa(solve(input, true)), nil
}

func part2(input string) (string, error) {
	return strconv.Itoa(solve(input, false)), nil
}

// Directions: 0=up, 1=right, 2=down, 3=left
var rightTurn = [4]int{1, 2, 3, 0}
var leftTurn = [4]int{3, 0, 1, 2}

type State struct {
	x, y, dir, score int
	prev             *State
}

func (s *State) nextX() int {
	if s.dir == 1 {
		return s.x + 1
	} else if s.dir == 3 {
		return s.x - 1
	}
	return s.x
}

func (s *State) nextY() int {
	if s.dir == 2 {
		return s.y + 1
	} else if s.dir == 0 {
		return s.y - 1
	}
	return s.y
}

func (s *State) coordId() int {
	return s.x*1000 + s.y
}

func (s *State) directedCoordId() int {
	return directedCoordId(s.x, s.y, s.dir)
}

func directedCoordId(x, y, dir int) int {
	return x*10000 + y*10 + dir
}

func (s *State) canGoForward(mapData [][]rune, visited map[int]int) bool {
	x2 := s.nextX()
	y2 := s.nextY()
	if y2 < 0 || y2 >= len(mapData) || x2 < 0 || x2 >= len(mapData[0]) {
		return false
	}
	if mapData[y2][x2] == '#' {
		return false
	}
	directed := directedCoordId(x2, y2, s.dir)
	visitedScore, ok := visited[directed]
	return !ok || visitedScore >= s.score+1
}

func (s *State) goForward() *State {
	return &State{s.nextX(), s.nextY(), s.dir, s.score + 1, s}
}

func (s *State) canTurnRight(visited map[int]int) bool {
	dir2 := rightTurn[s.dir]
	directed := directedCoordId(s.x, s.y, dir2)
	visitedScore, ok := visited[directed]
	return !ok || visitedScore >= s.score+1000
}

func (s *State) turnRight() *State {
	return &State{s.x, s.y, rightTurn[s.dir], s.score + 1000, s}
}

func (s *State) canTurnLeft(visited map[int]int) bool {
	dir2 := leftTurn[s.dir]
	directed := directedCoordId(s.x, s.y, dir2)
	visitedScore, ok := visited[directed]
	return !ok || visitedScore >= s.score+1000
}

func (s *State) turnLeft() *State {
	return &State{s.x, s.y, leftTurn[s.dir], s.score + 1000, s}
}

// Priority queue for State

type StatePQ []*State

func (pq StatePQ) Len() int           { return len(pq) }
func (pq StatePQ) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq StatePQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *StatePQ) Push(x interface{}) { *pq = append(*pq, x.(*State)) }
func (pq *StatePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func solve(input string, returnFirstBestPath bool) int {
	lines := utils.Lines(input)
	mapData := make([][]rune, len(lines))
	for i, line := range lines {
		mapData[i] = []rune(line)
	}

	paths := &StatePQ{}
	heap.Init(paths)
	visited := make(map[int]int, len(input))
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[0]); x++ {
			if mapData[y][x] == 'S' {
				start := &State{x, y, 1, 0, nil}
				heap.Push(paths, start)
				visited[start.directedCoordId()] = 0
				break
			}
		}
	}

	bestScore := -1
	allBestPathsVisited := make(map[int]struct{})
	for paths.Len() > 0 {
		s := heap.Pop(paths).(*State)
		if mapData[s.y][s.x] == 'E' {
			if bestScore == -1 || bestScore == s.score {
				bestScore = s.score
				if returnFirstBestPath {
					return s.score
				} else {
					curr := s
					for curr != nil {
						allBestPathsVisited[curr.coordId()] = struct{}{}
						curr = curr.prev
					}
				}
			}
			continue
		}
		if s.canGoForward(mapData, visited) {
			processState(s.goForward(), visited, paths)
		}
		if s.canTurnRight(visited) {
			processState(s.turnRight(), visited, paths)
		}
		if s.canTurnLeft(visited) {
			processState(s.turnLeft(), visited, paths)
		}
	}
	return len(allBestPathsVisited)
}

func processState(s2 *State, visited map[int]int, paths *StatePQ) {
	heap.Push(paths, s2)
	visited[s2.directedCoordId()] = s2.score
}
