package registry

import (
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

var solutions []*utils.Solution

func AddSolution(year int, day int, name string, input string, part1 func(string) string, part2 func(string) string) {
	solutions = append(solutions, &utils.Solution{
		Year:  year,
		Day:   day,
		Name:  name,
		Input: input,
		Part1: part1,
		Part2: part2,
	})
}

func GetSolutions() []*utils.Solution {
	return solutions
}
