package year2025day08

import (
	"cmp"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day08.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/8
	registry.AddSolution(2025, 8, "Playground", input, part1, part2)
}

func part1(input string) (string, error) {
	var result int
	lines := utils.Lines(input)
	n := len(lines)

	junctions := make([]junction, n)
	junctionToCircuit := make([]int, n)
	circuits := make(map[int][]int, n)
	err := initialiseCircuits(lines, junctions, junctionToCircuit, circuits)
	if err != nil {
		return "", utils.ErrBadInput
	}

	distances := make([]distance, 0, n*(n-1)/2)
	distances = prepareDistances(junctions, distances)

	var connections int
	if n == 20 {
		connections = 10 // sample
	} else {
		connections = 1000 // input
	}
	connectJunctions(connections, distances, junctionToCircuit, circuits)

	circuitsSizes := make([]int, 0, len(circuits))
	for _, c := range circuits {
		circuitsSizes = append(circuitsSizes, len(c))
	}

	slices.Sort(circuitsSizes)
	slices.Reverse(circuitsSizes)

	result = circuitsSizes[0] * circuitsSizes[1] * circuitsSizes[2]

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	var result int
	lines := utils.Lines(input)
	n := len(lines)

	junctions := make([]junction, n)
	junctionToCircuit := make([]int, n)
	circuits := make(map[int][]int, n)
	err := initialiseCircuits(lines, junctions, junctionToCircuit, circuits)
	if err != nil {
		return "", utils.ErrBadInput
	}

	distances := make([]distance, 0, n*(n-1)/2)
	distances = prepareDistances(junctions, distances)

	last := connectJunctions(len(distances), distances, junctionToCircuit, circuits)
	result = int(junctions[last.j1].x) * int(junctions[last.j2].x)

	return strconv.Itoa(result), nil
}

type junction struct {
	x, y, z int64
}

type distance struct {
	len    int64
	j1, j2 int
}

func initialiseCircuits(lines []string, junctions []junction, junctionToCircuit []int, circuits map[int][]int) error {
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return err
		}
		z, err := strconv.ParseInt(split[2], 10, 64)
		if err != nil {
			return err
		}
		junctions[i] = junction{x, y, z}
		junctionToCircuit[i] = i
		circuits[i] = []int{i}
	}
	return nil
}

func prepareDistances(junctions []junction, distances []distance) []distance {
	for i := 0; i < len(junctions)-1; i++ {
		for j := i + 1; j < len(junctions); j++ {
			distances = append(distances, distance{calculateDistance(junctions[i], junctions[j]), i, j})
		}
	}
	slices.SortFunc(distances, func(d1, d2 distance) int {
		return cmp.Compare(d1.len, d2.len)
	})
	return distances
}

func calculateDistance(j1, j2 junction) int64 {
	dx := abs(j1.x - j2.x)
	dy := abs(j1.y - j2.y)
	dz := abs(j1.z - j2.z)
	return dx*dx + dy*dy + dz*dz
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func connectJunctions(connections int, distances []distance, junctionToCircuit []int, circuits map[int][]int) distance {
	for i := range connections {
		dist := distances[i]
		c1 := junctionToCircuit[dist.j1]
		c2 := junctionToCircuit[dist.j2]

		if c1 == c2 {
			continue
		}

		circuits[c1] = append(circuits[c1], circuits[c2]...)
		for _, j := range circuits[c2] {
			junctionToCircuit[j] = c1
		}
		delete(circuits, c2)
		if len(circuits) == 1 {
			return dist
		}
	}
	return distance{}
}
