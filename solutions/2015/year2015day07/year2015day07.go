package year2015day07

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2015day07.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/7
	registry.AddSolution(2015, 7, "Some Assembly Required", input, part1, part2)
}

func part1(input string) (string, error) {
	gates, memo, err := prepareGates(input)
	if err != nil {
		return "", err
	}
	result, err := gateValue("a", gates, memo)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(result)), nil
}

func part2(input string) (string, error) {
	gates, memo, err := prepareGates(input)
	if err != nil {
		return "", err
	}
	signalA, err := gateValue("a", gates, memo)
	if err != nil {
		return "", err
	}
	clear(memo)
	memo["b"] = signalA
	result, err := gateValue("a", gates, memo)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(result)), nil
}

func gateValue(wire string, gates map[string]func() (uint16, error), memo map[string]uint16) (uint16, error) {
	if value, exists := memo[wire]; exists {
		return value, nil
	}

	var result uint16
	var err error
	if wire[0] > '9' {
		result, err = gates[wire]()
		if err != nil {
			return 0, err
		}
	} else {
		wire, err := strconv.Atoi(wire)
		if err != nil {
			return 0, err
		}
		result = uint16(wire)
	}

	memo[wire] = result
	return result, nil
}

func prepareGates(input string) (map[string]func() (uint16, error), map[string]uint16, error) {
	lines := strings.Split(input, "\n")
	gates := make(map[string]func() (uint16, error))
	memo := make(map[string]uint16)
	gateValue := func(wire string) (uint16, error) {
		return gateValue(wire, gates, memo)
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, " ")

		if len(split) == 3 { // value
			to := split[2]
			from := split[0]
			gates[to] = func() (uint16, error) {
				value, err := gateValue(from)
				if err != nil {
					return 0, err
				}
				return value, nil
			}
		} else if len(split) == 4 { // NOT
			to := split[3]
			from := split[1]
			gates[to] = func() (uint16, error) {
				value, err := gateValue(from)
				if err != nil {
					return 0, err
				}
				return ^value, nil
			}
		} else if len(split) == 5 { // AND, OR, LSHIFT, RSHIFT
			to := split[4]
			from1, from2 := split[0], split[2]
			switch op := split[1]; op {
			case "AND":
				gates[to] = func() (uint16, error) {
					value1, err := gateValue(from1)
					if err != nil {
						return 0, err
					}
					value2, err := gateValue(from2)
					if err != nil {
						return 0, err
					}
					return value1 & value2, nil
				}
			case "OR":
				gates[to] = func() (uint16, error) {
					value1, err := gateValue(from1)
					if err != nil {
						return 0, err
					}
					value2, err := gateValue(from2)
					if err != nil {
						return 0, err
					}
					return value1 | value2, nil
				}
			case "LSHIFT":
				gates[to] = func() (uint16, error) {
					value1, err := gateValue(from1)
					if err != nil {
						return 0, err
					}
					value2, err := gateValue(from2)
					if err != nil {
						return 0, err
					}
					return value1 << value2, nil
				}
			case "RSHIFT":
				gates[to] = func() (uint16, error) {
					value1, err := gateValue(from1)
					if err != nil {
						return 0, err
					}
					value2, err := gateValue(from2)
					if err != nil {
						return 0, err
					}
					return value1 >> value2, nil
				}
			}
		}

	}

	return gates, memo, nil
}
