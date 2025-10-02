package puzzles

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed year2015day07.txt
var year2015Day07Input string

func init() {
	// https://adventofcode.com/2015/day/7
	addPuzzle(2015, 7, "Some Assembly Required", year2015Day07Input, year2015Day07Part1, year2015Day07Part2)
}

func year2015Day07Part1(input string) string {
	gates, memo := prepareGates(input)
	result := int(gateValue("a", gates, memo))
	return strconv.Itoa(result)
}

func year2015Day07Part2(input string) string {
	gates, memo := prepareGates(input)
	signalA := gateValue("a", gates, memo)
	clear(memo)
	memo["b"] = signalA
	result := int(gateValue("a", gates, memo))
	return strconv.Itoa(result)
}

func gateValue(wire string, gates map[string]func() uint16, memo map[string]uint16) uint16 {
	if value, exists := memo[wire]; exists {
		return value
	}

	var result uint16
	if wire[0] > '9' {
		result = gates[wire]()
	} else {
		tmp, _ := strconv.Atoi(wire)
		result = uint16(tmp)
	}

	memo[wire] = result
	return result
}

func prepareGates(input string) (map[string]func() uint16, map[string]uint16) {
	lines := strings.Split(input, "\n")
	gates := make(map[string]func() uint16)
	memo := make(map[string]uint16)
	gateValue := func(wire string) uint16 {
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
			gates[to] = func() uint16 {
				value := gateValue(from)
				return value
			}
		} else if len(split) == 4 { // NOT
			to := split[3]
			from := split[1]
			gates[to] = func() uint16 {
				value := ^gateValue(from)
				//print(to, " = ", value, "\n")
				return value
			}
		} else if len(split) == 5 { // AND, OR, LSHIFT, RSHIFT
			to := split[4]
			from1, from2 := split[0], split[2]
			switch op := split[1]; op {
			case "AND":
				gates[to] = func() uint16 { return gateValue(from1) & gateValue(from2) }
			case "OR":
				gates[to] = func() uint16 { return gateValue(from1) | gateValue(from2) }
			case "LSHIFT":
				gates[to] = func() uint16 { return gateValue(from1) << gateValue(from2) }
			case "RSHIFT":
				gates[to] = func() uint16 { return gateValue(from1) >> gateValue(from2) }
			}
		}

	}

	return gates, memo
}
