package year2025day10

import (
	"container/list"
	_ "embed"
	"math"
	"slices"
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
	var result int

	machines, err := parseInput(input)
	if err != nil {
		return "", utils.ErrBadInput
	}

	for _, m := range machines {
		result += configureJoltage(&m)
	}

	return strconv.Itoa(result), nil
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

const epsilon = 1e-9
const maxInt = int(^uint(0) >> 1)

type matrix struct {
	data         [][]float64
	rows         int
	cols         int
	dependents   []int
	independents []int
}

func newMatrix(m *machine) *matrix {
	rows := len(m.joltage)
	cols := len(m.buttons)
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols+1)
	}

	for c, button := range m.buttons {
		for _, r := range button {
			if r < rows {
				data[r][c] = 1.0
			}
		}
	}

	for r, val := range m.joltage {
		data[r][cols] = float64(val)
	}

	mat := &matrix{
		data:         data,
		rows:         rows,
		cols:         cols,
		dependents:   make([]int, 0),
		independents: make([]int, 0),
	}

	mat.gaussianElimination()
	return mat
}

func (m *matrix) gaussianElimination() {
	pivot := 0
	col := 0

	for pivot < m.rows && col < m.cols {
		// Find the best pivot row for this column
		bestRow := pivot
		bestValue := math.Abs(m.data[pivot][col])
		for r := pivot + 1; r < m.rows; r++ {
			val := math.Abs(m.data[r][col])
			if val > bestValue {
				bestRow = r
				bestValue = val
			}
		}

		// If the best value is zero, this is a free variable
		if bestValue < epsilon {
			m.independents = append(m.independents, col)
			col++
			continue
		}

		// Swap rows and mark this column as dependent
		m.data[pivot], m.data[bestRow] = m.data[bestRow], m.data[pivot]
		m.dependents = append(m.dependents, col)

		// Normalize pivot row
		pivotValue := m.data[pivot][col]
		for c := col; c <= m.cols; c++ {
			m.data[pivot][c] /= pivotValue
		}

		// Eliminate this column in all other rows
		for r := 0; r < m.rows; r++ {
			if r != pivot {
				factor := m.data[r][col]
				if math.Abs(factor) > epsilon {
					for c := col; c <= m.cols; c++ {
						m.data[r][c] -= factor * m.data[pivot][c]
					}
				}
			}
		}

		pivot++
		col++
	}

	// Any remaining columns are free variables
	for ; col < m.cols; col++ {
		m.independents = append(m.independents, col)
	}
}

func (m *matrix) valid(values []int) (int, bool) {
	// Start with how many times we've pressed the free variables
	total := 0
	for _, v := range values {
		total += v
	}

	// Calculate dependent variable values based on independent variables
	for row := 0; row < len(m.dependents); row++ {
		val := m.data[row][m.cols]
		for i, col := range m.independents {
			val -= m.data[row][col] * float64(values[i])
		}

		// We need non-negative, whole numbers for a valid solution
		if val < -epsilon {
			return 0, false
		}
		rounded := math.Round(val)
		if math.Abs(val-rounded) > epsilon {
			return 0, false
		}

		total += int(rounded)
	}

	return total, true
}

func dfs(mat *matrix, idx int, values []int, minVal int, max int) int {
	result := minVal
	// When we've assigned all independent variables, check if it's a valid solution
	if idx == len(mat.independents) {
		if total, ok := mat.valid(values); ok {
			result = min(result, total)
		}
		return result
	}

	// Try different values for the current independent variable
	total := 0
	for i := 0; i < idx; i++ {
		total += values[i]
	}

	for val := 0; val < max; val++ {
		// Optimization: If we ever go above our min result, we can't possibly do better
		if total+val >= result {
			break
		}
		values[idx] = val
		result = dfs(mat, idx+1, values, result, max)
	}

	return result
}

func configureJoltage(m *machine) int {
	mat := newMatrix(m)
	maxVal := slices.Max(m.joltage) + 1
	values := make([]int, len(mat.independents))
	result := dfs(mat, 0, values, maxInt, maxVal)
	return result
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
