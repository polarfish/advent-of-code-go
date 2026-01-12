package year2024day21

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2024day21.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/21
	registry.AddSolution(2024, 21, "Keypad Conundrum", input, part1, part2)
}

func part1(input string) (string, error) {
	memo := make(map[MemoKey]int64)
	var result int64 = 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		size := translate(line, 2, 0, memo)
		codeNum, _ := strconv.Atoi(line[:3])
		result += size * int64(codeNum)
	}
	return strconv.FormatInt(result, 10), nil
}

func part2(input string) (string, error) {
	memo := make(map[MemoKey]int64)
	var result int64 = 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		size := translate(line, 25, 0, memo)
		codeNum, _ := strconv.Atoi(line[:3])
		result += size * int64(codeNum)
	}
	return strconv.FormatInt(result, 10), nil
}

type MemoKey struct {
	seq   string
	level int
}

var nPad = map[rune][][]rune{
	'0': {{'2', '^'}, {'A', '>'}},
	'1': {{'2', '>'}, {'4', '^'}},
	'2': {{'0', 'v'}, {'1', '<'}, {'3', '>'}, {'5', '^'}},
	'3': {{'2', '<'}, {'6', '^'}, {'A', 'v'}},
	'4': {{'1', 'v'}, {'5', '>'}, {'7', '^'}},
	'5': {{'2', 'v'}, {'4', '<'}, {'6', '>'}, {'8', '^'}},
	'6': {{'3', 'v'}, {'5', '<'}, {'9', '^'}},
	'7': {{'4', 'v'}, {'8', '>'}},
	'8': {{'5', 'v'}, {'7', '<'}, {'9', '>'}},
	'9': {{'6', 'v'}, {'8', '<'}},
	'A': {{'0', '<'}, {'3', '^'}},
}

var dPad = map[rune][][]rune{
	'^': {{'A', '>'}, {'v', 'v'}},
	'>': {{'A', '^'}, {'v', '<'}},
	'v': {{'^', '^'}, {'>', '>'}, {'<', '<'}},
	'<': {{'v', '>'}},
	'A': {{'>', 'v'}, {'^', '<'}},
}

var pads = []map[rune][][]rune{nPad, dPad}

func translate(seq string, level int, buttonsIndex int, memo map[MemoKey]int64) int64 {
	key := MemoKey{seq: seq, level: level}
	if v, ok := memo[key]; ok {
		return v
	}

	pad := pads[buttonsIndex]
	var result int64 = 0

	seq = "A" + seq

	pairs := make([]string, 0, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		pairs = append(pairs, seq[i:i+2])
	}

	for _, pair := range pairs {
		from := rune(pair[0])
		to := rune(pair[1])
		paths := search(from, to, pad)
		if level == 0 {
			// choose shortest path length
			min := -1
			for _, p := range paths {
				l := len(p)
				if min == -1 || l < min {
					min = l
				}
			}
			if min != -1 {
				result += int64(min)
			}
		} else {
			min := int64(-1)
			for _, p := range paths {
				v := translate(p, level-1, 1, memo)
				if min == -1 || v < min {
					min = v
				}
			}
			if min != -1 {
				result += min
			}
		}
	}

	memo[key] = result
	return result
}

type step struct {
	ch   rune
	path []rune
}

func search(from, to rune, pad map[rune][][]rune) []string {
	queue := make([]step, 0)
	queue = append(queue, step{ch: from, path: []rune{}})
	visited := make(map[rune]bool)
	visited[from] = true
	res := make([]string, 0)
	shortest := -1

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		curr := s.ch
		path := s.path
		if curr == to {
			if shortest == -1 {
				shortest = len(path)
			}
			if len(path) == shortest {
				// append 'A' as in Java solution
				out := string(path) + "A"
				res = append(res, out)
			}
			continue
		}
		if shortest != -1 && len(path) >= shortest {
			continue
		}
		for _, transitions := range pad[curr] {
			btn := transitions[0]
			dir := transitions[1]
			if visited[btn] {
				continue
			}
			path2 := make([]rune, len(path))
			copy(path2, path)
			path2 = append(path2, dir)
			queue = append(queue, step{ch: btn, path: path2})
		}
		visited[curr] = true
	}
	return res
}
