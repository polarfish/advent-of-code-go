package year2024day14

import (
	_ "embed"
	"regexp"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day14.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/14
	registry.AddSolution(2024, 14, "Restroom Redoubt", input, part1, part2)
}

var robotPattern = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

const (
	PX = iota
	PY
	VX
	VY
)

func part1(input string) (string, error) {
	robots, err := parseRobots(input)
	if err != nil {
		return "", err
	}
	w, h := 11, 7
	if len(robots) >= 100 {
		w, h = 101, 103
	}
	for i := range robots {
		moveRobot(robots[i], w, h, 100)
	}
	result := calculateSafetyFactor(robots, w, h)
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	robots, err := parseRobots(input)
	if err != nil {
		return "", err
	}
	w, h := 11, 7
	if len(robots) >= 100 {
		w, h = 101, 103
	}
	initialSafetyFactor := calculateSafetyFactor(robots, w, h)
	safetyFactor := initialSafetyFactor
	s := 0
	limit := 100_000
	for s < limit {
		s++
		for i := range robots {
			moveRobot(robots[i], w, h, 1)
		}
		safetyFactor = calculateSafetyFactor(robots, w, h)
		if abs(initialSafetyFactor-safetyFactor) >= 130_000_000 {
			break
		}
	}

	if s == limit {
		return "", utils.ErrIterSafetyLimit
	}

	return strconv.Itoa(s), nil
}

func parseRobots(input string) ([][]int, error) {
	var robots [][]int
	for _, line := range utils.Lines(input) {
		m := robotPattern.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		r := make([]int, 4)
		for i := 1; i <= 4; i++ {
			n, err := strconv.Atoi(m[i])
			if err != nil {
				return nil, utils.ErrBadInput
			}
			r[i-1] = n
		}
		robots = append(robots, r)
	}
	return robots, nil
}

func moveRobot(r []int, w, h, seconds int) {
	r[PX] = ((r[PX]+r[VX]*seconds)%w + w) % w
	r[PY] = ((r[PY]+r[VY]*seconds)%h + h) % h
}

func calculateSafetyFactor(robots [][]int, w, h int) int {
	quadrants := make([]int, 5)
	for _, r := range robots {
		var index int
		if r[PY] == h/2 || r[PX] == w/2 {
			index = 0
		} else if r[PY] < h/2 {
			if r[PX] < w/2 {
				index = 1
			} else {
				index = 2
			}
		} else {
			if r[PX] < w/2 {
				index = 3
			} else {
				index = 4
			}
		}
		quadrants[index]++
	}
	return quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
