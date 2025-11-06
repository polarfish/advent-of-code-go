package year2024day10

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day10.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/10
	registry.AddSolution(2024, 10, "Hoof It", input, part1, part2)
}

func part1(input string) (string, error) {
	mapData := parseMap(input)
	result := 0
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[0]); x++ {
			if mapData[y][x] == 0 {
				visitedPeaks := make(map[int]struct{})
				result += calculateScore(mapData, x, y, visitedPeaks)
			}
		}
	}
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {
	mapData := parseMap(input)
	result := 0
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[0]); x++ {
			if mapData[y][x] == 0 {
				result += calculateScore(mapData, x, y, nil)
			}
		}
	}
	return strconv.Itoa(result), nil
}

func parseMap(input string) [][]int {
	lines := utils.Lines(input)
	mapData := make([][]int, len(lines))
	for y, line := range lines {
		row := make([]int, len(line))
		for x, ch := range line {
			row[x] = int(ch - '0')
		}
		mapData[y] = row
	}
	return mapData
}

func calculateScore(mapData [][]int, x, y int, visitedPeaks map[int]struct{}) int {
	if mapData[y][x] == 9 {
		if visitedPeaks == nil {
			return 1
		}
		coordId := toCoordinatesId(x, y)
		if _, ok := visitedPeaks[coordId]; !ok {
			visitedPeaks[coordId] = struct{}{}
			return 1
		}
		return 0
	}
	result := 0
	nextHeight := mapData[y][x] + 1
	// up
	if y > 0 && mapData[y-1][x] == nextHeight {
		result += calculateScore(mapData, x, y-1, visitedPeaks)
	}
	// right
	if x < len(mapData[0])-1 && mapData[y][x+1] == nextHeight {
		result += calculateScore(mapData, x+1, y, visitedPeaks)
	}
	// down
	if y < len(mapData)-1 && mapData[y+1][x] == nextHeight {
		result += calculateScore(mapData, x, y+1, visitedPeaks)
	}
	// left
	if x > 0 && mapData[y][x-1] == nextHeight {
		result += calculateScore(mapData, x-1, y, visitedPeaks)
	}
	return result
}

func toCoordinatesId(x, y int) int {
	return x*100 + y
}
