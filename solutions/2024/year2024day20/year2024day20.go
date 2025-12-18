package year2024day20

import (
    _ "embed"
    "strconv"

    "github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2024day20.txt
var input string

func init() {
    // https://adventofcode.com/2024/day/20
    registry.AddSolution(2024, 20, "Race Condition", input, part1, part2)
}

func part1(input string) (string, error) {
    return strconv.Itoa(0), nil
}

func part2(input string) (string, error) {
    return strconv.Itoa(0), nil
}
