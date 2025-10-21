package utils

import (
	"bufio"
	"errors"
	"strings"
	"time"
)

var ErrBadInput = errors.New("bad input")
var ErrIterSafetyLimit = errors.New("reached iterations safety limit")

func Lines(input string) []string {
	var lines []string
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type Solution struct {
	Day   int
	Year  int
	Name  string
	Input string
	Part1 func(input string) (string, error)
	Part2 func(input string) (string, error)
}

type Result struct {
	Solution  *Solution
	Result1   string
	Error1    error
	Result2   string
	Error2    error
	Duration1 time.Duration
	Duration2 time.Duration
}

func (s Solution) Run() *Result {
	r := Result{Solution: &s}

	start1 := time.Now()
	r.Result1, r.Error1 = runSafe(s.Part1, s.Input)
	r.Duration1 = time.Since(start1)

	start2 := time.Now()
	r.Result2, r.Error2 = runSafe(s.Part2, s.Input)
	r.Duration2 = time.Since(start2)

	return &r
}

func runSafe(part func(input string) (string, error), input string) (res string, err error) {
	defer func() {
		if r := recover(); r != nil {
			res, err = "", errors.New("panic: "+r.(string))
		}
	}()
	res, err = part(input)
	return
}
