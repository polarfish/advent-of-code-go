package year2024day19

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day19.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/19
	registry.AddSolution(2024, 19, "Linen Layout", input, part1, part2)
}

func part1(input string) (string, error) {
	patterns, designs, err := parseInput(input)
	if err != nil {
		return "", utils.ErrBadInput
	}

	memo := make(map[string]int64)
	trie := NewTrieNodeWithPatterns(patterns)
	result := 0
	for _, design := range designs {
		if countCombinations(design, patterns, trie, memo, true) > 0 {
			result++
		}
	}

	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	patterns, designs, err := parseInput(input)
	if err != nil {
		return "", utils.ErrBadInput
	}

	var result int64 = 0
	memo := make(map[string]int64)
	trie := NewTrieNodeWithPatterns(patterns)
	for _, design := range designs {
		result += countCombinations(design, patterns, trie, memo, false)
	}

	return strconv.FormatInt(result, 10), nil
}

func countCombinations(design string, patterns []string, trie *TrieNode, memo map[string]int64, stopOnFirstResult bool) int64 {
	if design == "" {
		return 1
	}
	if val, ok := memo[design]; ok {
		return val
	}

	var result int64 = 0
	sizes := trie.FindStartPatternsSizes(design)
	if stopOnFirstResult {
		// it's faster to start with the longest patterns, when we stop on first result
		slices.Reverse(sizes)
	}
	for _, size := range sizes {
		combinations := countCombinations(design[size:], patterns, trie, memo, stopOnFirstResult)
		result += combinations
		if stopOnFirstResult && result > 0 {
			break
		}
	}

	memo[design] = result
	return result
}

func parseInput(input string) ([]string, []string, error) {
	lines := utils.Lines(input)
	if len(lines) < 2 {
		return nil, nil, utils.ErrBadInput
	}

	patterns := strings.Split(lines[0], ", ")
	designs := lines[2:]

	return patterns, designs, nil
}

type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func NewTrieNodeWithPatterns(patterns []string) *TrieNode {
	node := &TrieNode{}
	for _, p := range patterns {
		node.Insert(p)
	}
	return node
}

func (t *TrieNode) Insert(word string) {
	current := t

	for i := 0; i < len(word); i++ {
		ind := word[i] - 'a'
		if current.children[ind] == nil {
			current.children[ind] = &TrieNode{}
		}
		current = current.children[ind]
	}
	current.isWord = true
}

func (t *TrieNode) FindStartPatternsSizes(design string) []int {
	var result []int
	current := t
	size := 0

	for size < len(design) {
		ind := design[size] - 'a'
		if current.children[ind] == nil {
			break
		}
		current = current.children[ind]
		size++
		if current.isWord {
			result = append(result, size)
		}
	}

	return result
}
