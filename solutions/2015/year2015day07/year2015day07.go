package year2015day07

import (
	_ "embed"
	"strings"

	"github.com/polarfish/advent-of-code-go/registry"
	"github.com/polarfish/advent-of-code-go/utils"
)

//go:embed year2015day07.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/7
	registry.AddSolution(2015, 7, "Some Assembly Required", input, part1, part2)
}

func part1(input string) string {
	gates, memo := prepareGates(input)
	result := int(gateValue("a", gates, memo))
	return utils.ToStr(result)
}

func part2(input string) string {
	gates, memo := prepareGates(input)
	signalA := gateValue("a", gates, memo)
	clear(memo)
	memo["b"] = signalA
	result := int(gateValue("a", gates, memo))
	return utils.ToStr(result)
}

func gateValue(wire string, gates map[string]func() uint16, memo map[string]uint16) uint16 {
	if value, exists := memo[wire]; exists {
		return value
	}

	var result uint16
	if wire[0] > '9' {
		result = gates[wire]()
	} else {
		result = uint16(utils.ToInt(wire))
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
