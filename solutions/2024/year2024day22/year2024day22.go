package year2024day22

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day22.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/22
	registry.AddSolution(2024, 22, "Monkey Market", input, part1, part2)
}

func part1(input string) (string, error) {
	var sum int64 = 0

	for _, line := range utils.Lines(input) {
		var n int64
		// numbers fit in 64-bit
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return "", err
		}
		for range 2000 {
			n = transform(n)
		}
		sum += n
	}
	return strconv.FormatInt(sum, 10), nil
}

func part2(input string) (string, error) {
	seq := make([]int, 4)
	i := 0
	totals := make(map[marker]int)

	for _, line := range utils.Lines(input) {
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return "", err
		}
		n2 := n
		currents := make(map[marker]int, 3000)
		for range 2000 {
			n2 = transform(n2)

			p := int(n % 10)
			p2 := int(n2 % 10)
			seq[i] = p2 - p
			i++
			if i == 4 {
				m := marker{seq[0], seq[1], seq[2], seq[3]}
				if _, ok := currents[m]; !ok {
					currents[m] = p2
				}
				// remove first
				seq[0] = seq[1]
				seq[1] = seq[2]
				seq[2] = seq[3]
				i--
			}

			n = n2
		}

		for k, v := range currents {
			totals[k] += v
		}
	}

	max := 0
	for _, v := range totals {
		if v > max {
			max = v
		}
	}

	return strconv.Itoa(max), nil
}

func transform(n int64) int64 {
	n = (n ^ (n << 6)) & 16777215
	n = (n ^ (n >> 5)) & 16777215
	n = (n ^ (n << 11)) & 16777215
	return n
}

type marker struct {
	d1, d2, d3, d4 int
}
