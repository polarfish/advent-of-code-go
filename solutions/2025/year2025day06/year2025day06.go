package year2025day06

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day06.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/6
	registry.AddSolution(2025, 6, "Trash Compactor", input, part1, part2)
}

func part1(input string) (string, error) {
	var result int64

	lines := utils.Lines(input)

	nums := make([][]int64, len(lines)-1)
	for i := range len(nums) {
		fields := strings.Fields(lines[i])
		nums[i] = make([]int64, len(fields))
		for j, field := range fields {
			num, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				return "", utils.ErrBadInput
			}
			nums[i][j] = num
		}
	}

	fields := strings.Fields(lines[len(lines)-1])
	signs := make([]byte, len(fields))
	for j, field := range fields {
		signs[j] = field[0]
	}

	for i, sign := range signs {
		curr := nums[0][i]
		for j := 1; j < len(nums); j++ {
			switch sign {
			case '+':
				curr += nums[j][i]
			case '*':
				curr *= nums[j][i]
			}
		}
		result += curr
	}

	return strconv.FormatInt(result, 10), nil
}

func part2(input string) (string, error) {
	var result int64
	lines := utils.Lines(input)

	fields := strings.Fields(lines[len(lines)-1])
	signs := make([]byte, len(fields))
	for j, field := range fields {
		signs[j] = field[0]
	}

	j := 0
	for _, sign := range signs {
		var curr int64
		switch sign {
		case '+':
			curr = 0
		case '*':
			curr = 1
		}

		for ; j < len(lines[0]); j++ {
			var num int64
			shouldStop := true
			for k := 0; k < len(lines)-1; k++ {
				if lines[k][j] != ' ' {
					shouldStop = false
					num = num*10 + int64(lines[k][j]-'0')
				}
			}

			if shouldStop {
				j++
				break
			}

			switch sign {
			case '+':
				curr += num
			case '*':
				curr *= num
			}
		}

		result += curr
	}

	return strconv.FormatInt(result, 10), nil
}
