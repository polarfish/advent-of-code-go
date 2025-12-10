package year2025day10

import (
	"container/list"
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day10.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/10
	registry.AddSolution(2025, 10, "Factory", input, part1, part2)
}

func part1(input string) (string, error) {
	var result int

	machines, err := parseInput(input)
	if err != nil {
		return "", utils.ErrBadInput
	}

	for _, m := range machines {
		result += configureIndicators(&m)
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	return strconv.Itoa(0), nil
}

type machine struct {
	indicators int // stored as bitmask, max 10 indicators supported
	buttons    [][]int
	joltage    []int
}

type state struct {
	indicators, presses int
}

func configureIndicators(m *machine) int {
	visited := make([]int, 1024)
	visited[0] = 1
	q := list.New()
	q.PushBack(state{0, 0})

	for q.Len() > 0 {
		curr := q.Remove(q.Front()).(state)
		for _, button := range m.buttons {
			nextIndicators := curr.indicators
			for _, indicatorWire := range button {
				nextIndicators ^= 1 << indicatorWire
			}
			nextPresses := curr.presses + 1

			if nextIndicators == m.indicators {
				return nextPresses
			}

			if visited[nextIndicators] == 0 {
				visited[nextIndicators] = 1
				q.PushBack(state{nextIndicators, nextPresses})
			}
		}
	}

	return 0
}

func parseInput(input string) ([]machine, error) {
	lines := utils.Lines(input)
	machines := make([]machine, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")

		// parsing indicators
		splitFirst := split[0]
		indicatorsPart := splitFirst[1 : len(splitFirst)-1]
		indicators := 0
		for i, ch := range indicatorsPart {
			if ch == '#' {
				indicators |= 1 << i
			}
		}

		// parsing buttons
		buttonsParts := split[1 : len(split)-1]
		buttons := make([][]int, len(buttonsParts))
		for i, s := range buttonsParts {
			buttonPart := s[1 : len(s)-1]
			buttonPartSplit := strings.Split(buttonPart, ",")
			button := make([]int, len(buttonPartSplit))
			for j, s := range buttonPartSplit {
				buttonIndicator, err := strconv.Atoi(s)
				if err != nil {
					return nil, err
				}
				button[j] = buttonIndicator
			}
			buttons[i] = button
		}

		// parsing joltage
		splitN := split[len(split)-1]
		joltagePart := splitN[1 : len(splitN)-1]
		joltagePartSplit := strings.Split(joltagePart, ",")
		joltage := make([]int, len(joltagePartSplit))
		for i, s := range joltagePartSplit {
			jol, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			joltage[i] = jol
		}

		machines = append(machines, machine{indicators, buttons, joltage})
	}

	return machines, nil
}
