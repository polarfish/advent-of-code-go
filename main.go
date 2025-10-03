package main

import (
	"sort"

	_ "github.com/polarfish/advent-of-code-go/puzzles/loader"
	"github.com/polarfish/advent-of-code-go/puzzles/registry"
	"github.com/polarfish/advent-of-code-go/puzzles/utils"

	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"
	"time"
)

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
	if err := syscall.SetNonblock(fd, true); err != nil {
		fmt.Println("Failed to set non-blocking mode")
		os.Exit(1)
	}
	stdInputBytes, _ := io.ReadAll(os.Stdin)
	stdInput := string(stdInputBytes)

	var puzzlesToRun []*utils.Puzzle
	for _, p := range registry.GetAllPuzzles() {
		if (day == 0 || day == p.Day) && (year == 0 || year == p.Year) {
			puzzlesToRun = append(puzzlesToRun, p)
		}
	}

	if len(puzzlesToRun) == 0 {
		fmt.Println("Not found")
		os.Exit(0)
	}

	sort.Slice(puzzlesToRun, func(i, j int) bool {
		if puzzlesToRun[i].Year != puzzlesToRun[j].Year {
			return puzzlesToRun[i].Year < puzzlesToRun[j].Year
		} else {
			return puzzlesToRun[i].Day < puzzlesToRun[j].Day
		}
	})

	var totalStart time.Time
	var totalElapsed time.Duration
	results := make([]utils.Result, len(puzzlesToRun))

	// run
	totalStart = time.Now()
	for i, puzzle := range puzzlesToRun {
		if stdInput != "" {
			results[i] = puzzle.Run(stdInput)
		} else {
			results[i] = puzzle.Run(puzzle.Input)
		}
	}
	totalElapsed = time.Since(totalStart)

	// report
	for _, r := range results {
		fmt.Printf("--- %d Day %d: %s ---\n", r.Puzzle.Year, r.Puzzle.Day, r.Puzzle.Name)
		fmt.Println(formatDuration(r.Duration1), "Part 1:", r.Result1)
		fmt.Println(formatDuration(r.Duration2), "Part 2:", r.Result2)
		fmt.Println()
	}

	if len(puzzlesToRun) > 1 {
		fmt.Println("=== Total ===")
		fmt.Println(formatDuration(totalElapsed))
	}
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("[%d.%03d ms]", d.Microseconds()/1000, d.Microseconds()%1000)
}
