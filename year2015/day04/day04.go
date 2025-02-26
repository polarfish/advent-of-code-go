package year2015day04

import (
	"crypto/md5"
	_ "embed"
	"io"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/utils"
)

//go:embed day04.txt
var input string

func New() *utils.Puzzle {
	return &utils.Puzzle{
		Year:  2015,
		Day:   4,
		Name:  "The Ideal Stocking Stuffer",
		Input: input,
		Part1: Part1,
		Part2: Part2,
	}
}

func Part1(input string) string {
	res, err := solve(input, func(result []byte) bool {
		return result[0] == 0 && result[1] == 0 && result[2] < 16
	})
	if err != nil {
		return utils.ERR
	}
	return res
}

func Part2(input string) string {
	res, err := solve(input, func(result []byte) bool {
		return result[0] == 0 && result[1] == 0 && result[2] == 0
	})
	if err != nil {
		return utils.ERR
	}
	return res
}

func solve(input string, predicate func([]byte) bool) (string, error) {
	input = strings.TrimSpace(input)
	h := md5.New()
	var result []byte
	for i := 1; i < 100_000_000; i++ {
		h.Reset()
		_, err1 := io.WriteString(h, input)
		if err1 != nil {
			return utils.ERR, nil
		}
		_, err2 := io.WriteString(h, strconv.Itoa(i))
		if err2 != nil {
			return utils.ERR, nil
		}
		result = h.Sum(nil)

		if predicate(result) {
			return strconv.Itoa(i), nil
		}
	}
	return utils.ERR, nil
}
