package main

import (
	"sort"

	"fmt"
	"os"
	"time"

	_ "github.com/polarfish/advent-of-code-go/tools/loader"
	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

func main() {
	args := os.Args[1:]

	var year int
	if len(args) > 0 {
		year = utils.ToInt(args[0])
	}

	var day int
	if len(args) > 1 {
		day = utils.ToInt(args[1])
	}

	var solutionsToRun []*utils.Solution
	for _, s := range registry.GetSolutions() {
		if (day == 0 || day == s.Day) && (year == 0 || year == s.Year) {
			solutionsToRun = append(solutionsToRun, s)
		}
	}

	if len(solutionsToRun) == 0 {
		fmt.Println("Not found")
		os.Exit(0)
	}

	sort.Slice(solutionsToRun, func(i, j int) bool {
		if solutionsToRun[i].Year != solutionsToRun[j].Year {
			return solutionsToRun[i].Year < solutionsToRun[j].Year
		} else {
			return solutionsToRun[i].Day < solutionsToRun[j].Day
		}
	})

	var totalStart time.Time
	var totalElapsed time.Duration
	results := make([]utils.Result, len(solutionsToRun))

	// run
	totalStart = time.Now()
	for i, solution := range solutionsToRun {
		results[i] = solution.Run()
	}
	totalElapsed = time.Since(totalStart)

	// report
	for _, r := range results {
		fmt.Printf("--- %d Day %d: %s ---\n", r.Solution.Year, r.Solution.Day, r.Solution.Name)
		fmt.Println(formatDuration(r.Duration1), "Part 1:", r.Result1)
		fmt.Println(formatDuration(r.Duration2), "Part 2:", r.Result2)
		fmt.Println()
	}

	if len(solutionsToRun) > 1 {
		fmt.Println("=== Total ===")
		fmt.Println(formatDuration(totalElapsed), len(solutionsToRun), "days total")
	}
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("[%d.%03d ms]", d.Microseconds()/1000, d.Microseconds()%1000)
}
