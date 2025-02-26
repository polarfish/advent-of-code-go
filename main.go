package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/polarfish/advent-of-code-go/utils"
	"github.com/polarfish/advent-of-code-go/year2015/day01"
	"github.com/polarfish/advent-of-code-go/year2015/day02"
	"github.com/polarfish/advent-of-code-go/year2015/day03"
	"github.com/polarfish/advent-of-code-go/year2015/day04"
)

var allPuzzles = []*utils.Puzzle{
	year2015day01.New(),
	year2015day02.New(),
	year2015day03.New(),
	year2015day04.New(),
}

func main() {
	args := os.Args[1:]

	var year int
	if len(args) > 0 {
		year, _ = strconv.Atoi(args[0])
	}

	var day int
	if len(args) > 1 {
		day, _ = strconv.Atoi(args[1])
	}

	fd := int(os.Stdin.Fd())
	err := syscall.SetNonblock(fd, true)
	if err != nil {
		fmt.Println("Failed to set non-blocking mode")
		os.Exit(1)
	}
	stdInputBytes, _ := io.ReadAll(os.Stdin)
	stdInput := string(stdInputBytes)

	var puzzles []*utils.Puzzle
	for _, p := range allPuzzles {
		if (day == 0 || day == p.Day) && (year == 0 || year == p.Year) {
			puzzles = append(puzzles, p)
		}
	}

	if len(puzzles) == 0 {
		fmt.Println("Not found")
		os.Exit(0)
	}

	var start, totalStart time.Time
	var totalElapsed time.Duration
	results := make([]*utils.Result, len(puzzles))

	// run
	totalStart = time.Now()
	for i, p := range puzzles {
		result := utils.Result{Puzzle: p}

		var input string
		if len(stdInput) > 0 {
			input = stdInput
		} else {
			input = p.Input
		}

		start = time.Now()
		result.Result1 = p.Part1(input)
		result.Duration1 = time.Since(start)

		start = time.Now()
		result.Result2 = p.Part2(input)
		result.Duration2 = time.Since(start)

		results[i] = &result
	}
	totalElapsed = time.Since(totalStart)

	// report
	for _, r := range results {
		fmt.Printf("--- %d Day %d: %s ---\n", r.Puzzle.Year, r.Puzzle.Day, r.Puzzle.Name)
		fmt.Println(utils.FormatDuration(r.Duration1), "Part 1:", r.Result1)
		fmt.Println(utils.FormatDuration(r.Duration2), "Part 2:", r.Result2)
		fmt.Println()
	}

	if len(puzzles) > 1 {
		fmt.Println("=== Total ===")
		fmt.Println(utils.FormatDuration(totalElapsed))
	}
}
