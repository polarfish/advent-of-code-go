package utils

import (
	"strings"
	"testing"
	"time"
)

const ErrorResult string = "error"
const NaResult string = "n/a"

func Lines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func Abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}

func Test(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
	}
}

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
