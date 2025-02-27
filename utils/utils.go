package utils

import (
	"fmt"
	"time"
)

const ERR string = "error"

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

var allPuzzles []*Puzzle

func RegisterPuzzle(p *Puzzle) {
	allPuzzles = append(allPuzzles, p)
}

func GetAllPuzzles() []*Puzzle {
	return allPuzzles
}

func FormatDuration(d time.Duration) string {
	return fmt.Sprintf("[%d.%03d ms]", d.Microseconds()/1000, d.Microseconds()%1000)
}
