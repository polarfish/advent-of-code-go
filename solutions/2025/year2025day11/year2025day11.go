package year2025day11

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day11.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/11
	registry.AddSolution(2025, 11, "Reactor", input, part1, part2)
}

func part1(input string) (string, error) {
	devices := parseInput(input)
	memo := map[string]map[string]int{}
	seen := map[string]struct{}{}
	result := searchAllPaths(devices, "you", "out", seen, memo)
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	devices := parseInput(input)
	memo := map[string]map[string]int{}

	result := 0
	fftDac := searchAllPaths(devices, "fft", "dac", map[string]struct{}{}, memo)
	if fftDac == 0 { // svr -> dac -> fft -> out
		svrDac := searchAllPaths(devices, "svr", "dac", map[string]struct{}{"fft": {}}, memo)
		dacFft := searchAllPaths(devices, "dac", "fft", map[string]struct{}{}, memo)
		fftOut := searchAllPaths(devices, "fft", "out", map[string]struct{}{"dac": {}}, memo)
		result = svrDac * dacFft * fftOut
	} else { // svr -> fft -> dac -> out
		svrFft := searchAllPaths(devices, "svr", "fft", map[string]struct{}{"dac": {}}, memo)
		dacOut := searchAllPaths(devices, "dac", "out", map[string]struct{}{"fft": {}}, memo)
		result = svrFft * fftDac * dacOut
	}

	return strconv.Itoa(result), nil
}

func searchAllPaths(devices map[string][]string, from, to string, seen map[string]struct{}, memo map[string]map[string]int) int {
	if from == to {
		return 1
	}

	cache, ok := memo[from]
	if !ok {
		cache = make(map[string]int)
		memo[from] = cache
	}

	if val, ok := cache[to]; ok {
		return val
	}

	seen[from] = struct{}{}

	totalPaths := 0
	for _, connection := range devices[from] {
		if _, ok := seen[connection]; !ok {
			paths2 := searchAllPaths(devices, connection, to, seen, memo)
			totalPaths += paths2
		}
	}
	delete(seen, from)
	cache[to] = totalPaths
	return totalPaths
}

func parseInput(input string) map[string][]string {
	devices := make(map[string][]string)
	for _, line := range utils.Lines(input) {
		split := strings.Split(line, ": ")
		in := split[0]
		out := strings.Split(split[1], " ")
		devices[in] = out
	}
	return devices
}
