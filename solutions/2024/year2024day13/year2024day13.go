package year2024day13

import (
	_ "embed"
	"regexp"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2024day13.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/13
	registry.AddSolution(2024, 13, "Claw Contraption", input, part1, part2)
}

var machinePattern = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

const (
	ax = iota
	ay
	bx
	by
	px
	py
)

func part1(input string) (string, error) {
	machines := parseMachines(input)
	var result int64
	for _, m := range machines {
		result += solveMachine(m)
	}
	return strconv.FormatInt(result, 10), nil
}

func part2(input string) (string, error) {
	machines := parseMachines(input)
	var result int64
	for _, m := range machines {
		m[px] += 10000000000000
		m[py] += 10000000000000
		result += solveMachine(m)
	}
	return strconv.FormatInt(result, 10), nil
}

func parseMachines(input string) [][]int64 {
	matches := machinePattern.FindAllStringSubmatch(input, -1)
	machines := make([][]int64, 0, len(matches))
	for _, m := range matches {
		machine := make([]int64, 6)
		for i := 1; i <= 6; i++ {
			val, _ := strconv.ParseInt(m[i], 10, 64)
			machine[i-1] = val
		}
		machines = append(machines, machine)
	}
	return machines
}

func solveMachine(m []int64) int64 {
	det := m[ax]*m[by] - m[bx]*m[ay]
	if det == 0 {
		return 0
	}
	a := (m[px]*m[by] - m[py]*m[bx]) / det
	b := (m[py]*m[ax] - m[px]*m[ay]) / det
	if m[ax]*a+m[bx]*b == m[px] && m[ay]*a+m[by]*b == m[py] {
		return a*3 + b
	}
	return 0
}
