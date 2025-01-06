package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/polarfish/advent-of-code-go/utils"
	year2015day01 "github.com/polarfish/advent-of-code-go/year2015/day01"
)

var allPuzzles = []*utils.Puzzle{
	year2015day01.New(),
}

func main() {
	args := os.Args[1:]

	year := -1
	if len(args) > 0 {
		year, _ = strconv.Atoi(args[0])
	}

	day := -1
	if len(args) > 1 {
		day, _ = strconv.Atoi(args[1])
	}

	fd := int(os.Stdin.Fd())
	syscall.SetNonblock(fd, true)
	stdInputBytes, _ := io.ReadAll(os.Stdin)
	stdInput := string(stdInputBytes)

	puzzles := []*utils.Puzzle{}
	for _, p := range allPuzzles {
		if (day == -1 || day == p.Day) && (year == -1 || year == p.Year) {
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

		var input *string
		if len(stdInput) > 0 {
			input = &stdInput
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
	fmt.Println("Total")
	fmt.Println(utils.FormatDuration(totalElapsed))
}
