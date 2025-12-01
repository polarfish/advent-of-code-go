package year2024day17

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day17.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/17
	registry.AddSolution(2024, 17, "Chronospatial Computer", input, part1, part2)
}

var programPattern = regexp.MustCompile(`Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: (.+)`)

const (
	A = 0
	B = 1
	C = 2
)

func part1(input string) (string, error) {
	matches := programPattern.FindStringSubmatch(input)
	if matches == nil {
		return "", utils.ErrBadInput
	}

	registers := make([]int64, 3)
	var err error
	registers[A], err = strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return "", utils.ErrBadInput
	}
	registers[B], err = strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return "", utils.ErrBadInput
	}
	registers[C], err = strconv.ParseInt(matches[3], 10, 64)
	if err != nil {
		return "", utils.ErrBadInput
	}

	programStr := strings.Split(matches[4], ",")
	program := make([]int, len(programStr))
	for i, s := range programStr {
		program[i], err = strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return "", utils.ErrBadInput
		}
	}

	out := make([]int, 32)
	outLen := runProgram(program, registers, out)

	result := make([]string, outLen)
	for i := 0; i < outLen; i++ {
		result[i] = strconv.Itoa(out[i])
	}

	return strings.Join(result, ","), nil
}

func part2(input string) (string, error) {
	matches := programPattern.FindStringSubmatch(input)
	if matches == nil {
		return "", utils.ErrBadInput
	}

	programStr := strings.Split(matches[4], ",")
	program := make([]int, len(programStr))
	var err error
	for i, s := range programStr {
		program[i], err = strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return "", utils.ErrBadInput
		}
	}

	registers := make([]int64, 3)
	out := make([]int, len(program))
	result := backtrack(program, registers, out, 0, len(program)-1)

	return strconv.FormatInt(result, 10), nil
}

func backtrack(program []int, registers []int64, out []int, a int64, dig int) int64 {
	res := int64(-1)

	for i := 0; i < 8; i++ {
		registers[A] = a + int64(i)
		registers[B] = 0
		registers[C] = 0
		runProgram(program, registers, out)
		if out[0] == program[dig] {
			if dig == 0 {
				return a + int64(i)
			}

			res2 := backtrack(program, registers, out, (a+int64(i))*8, dig-1)
			if res2 != -1 {
				if res == -1 {
					res = res2
				} else if res2 < res {
					res = res2
				}
			}
		}
	}

	return res
}

func runProgram(program []int, registers []int64, out []int) int {
	outLen := 0
	ptr := 0
	for ptr < len(program) {
		opcode := program[ptr]
		ptr++
		operand := program[ptr]
		ptr++

		switch opcode {
		case 0: // adv
			registers[A] = registers[A] / pow(2, combo(registers, operand), 1)
		case 1: // bxl
			registers[B] = registers[B] ^ int64(operand)
		case 2: // bst
			registers[B] = combo(registers, operand) % 8
		case 3: // jnz
			if registers[A] != 0 {
				ptr = operand
			}
		case 4: // bxc
			registers[B] = registers[B] ^ registers[C]
		case 5: // out
			out[outLen] = int(combo(registers, operand) % 8)
			outLen++
		case 6: // bdv
			registers[B] = registers[A] / pow(2, combo(registers, operand), 1)
		case 7: // cdv
			registers[C] = registers[A] / pow(2, combo(registers, operand), 1)
		}
	}
	return outLen
}

func combo(registers []int64, operand int) int64 {
	if operand > 3 && operand < 7 {
		return registers[operand-4]
	}
	return int64(operand)
}

func pow(n, p, accumulator int64) int64 {
	if p == 0 {
		return accumulator
	}
	return pow(n, p-1, accumulator*n)
}
