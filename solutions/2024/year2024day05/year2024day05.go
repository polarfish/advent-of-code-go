package year2024day05

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day05.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/5
	registry.AddSolution(2024, 5, "Print Queue", input, part1, part2)
}

func part1(input string) (string, error) {
	rules, updates, err := extractInput(input)
	if err != nil {
		return "", err
	}
	cmp := createCmpFunc(rules)
	result := 0
	for _, pages := range updates {
		if slices.IsSortedFunc(pages, cmp) {
			result += getMiddlePage(pages)
		}
	}
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	rules, updates, err := extractInput(input)
	if err != nil {
		return "", err
	}
	cmp := createCmpFunc(rules)
	result := 0
	for _, pages := range updates {
		if slices.IsSortedFunc(pages, cmp) {
			continue
		}
		slices.SortFunc(pages, cmp)
		result += getMiddlePage(pages)
	}
	return strconv.Itoa(result), nil
}

func extractInput(input string) (map[int]struct{}, [][]int, error) {
	rules := make(map[int]struct{})
	var updates [][]int
	rulesFinished := false
	for _, line := range utils.Lines(input) {
		if line == "" {
			rulesFinished = true
			continue
		}
		if !rulesFinished {
			split := strings.Split(line, "|")
			var err error

			before, err := strconv.Atoi(split[0])
			if err != nil {
				return nil, nil, utils.ErrBadInput
			}

			after, err := strconv.Atoi(split[1])
			if err != nil {
				return nil, nil, utils.ErrBadInput
			}

			rules[createRuleId(before, after)] = struct{}{}
		} else {
			split := strings.Split(line, ",")
			pages := make([]int, len(split))
			for i, s := range split {
				num, err := strconv.Atoi(s)
				if err != nil {
					return nil, nil, utils.ErrBadInput
				}
				pages[i] = num
			}
			updates = append(updates, pages)
		}
	}
	return rules, updates, nil
}

func createCmpFunc(rules map[int]struct{}) func(a, b int) int {
	return func(a, b int) int {
		if _, ok := rules[createRuleId(a, b)]; ok {
			return -1
		}
		return 1
	}
}

func createRuleId(beforePage, afterPage int) int {
	return beforePage*100 + afterPage
}

func getMiddlePage(pages []int) int {
	return pages[len(pages)/2]
}
