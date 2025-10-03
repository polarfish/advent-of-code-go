package utils

import (
	"time"
)

const ErrorResult string = "error"
const NaResult string = "n/a"

type Puzzle struct {
	Day   int
	Year  int
	Name  string
	Input string
	Part1 func(input string) string
	Part2 func(input string) string
}

type Result struct {
	Puzzle    *Puzzle
	Result1   string
	Result2   string
	Duration1 time.Duration
	Duration2 time.Duration
}

func (p Puzzle) Run() Result {
	result := Result{Puzzle: &p}

	start1 := time.Now()
	result.Result1 = runSafe(p.Part1, p.Input)
	result.Duration1 = time.Since(start1)

	start2 := time.Now()
	result.Result2 = runSafe(p.Part2, p.Input)
	result.Duration2 = time.Since(start2)

	return result
}

func runSafe(part func(input string) string, input string) (result string) {
	defer func() {
		if r := recover(); r != nil {
			result = ErrorResult
		}
	}()
	result = part(input)
	return result
}
