package year2024day12

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2024day12.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/12
	registry.AddSolution(2024, 12, "Garden Groups", input, part1, part2)
}

const (
	area = iota
	perimeter
	sides
)

func part1(input string) (string, error) {
	mapData := parseMap(input)
	visited := make([][]bool, len(mapData))
	for i := range visited {
		visited[i] = make([]bool, len(mapData[0]))
	}
	result := 0
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[0]); x++ {
			if region := measureRegion(mapData, x, y, visited); region != nil {
				result += region[area] * region[perimeter]
			}
		}
	}
	return strconv.Itoa(result), nil
}

func part2(input string) (string, error) {

	mapData := parseMap(input)
	visited := make([][]bool, len(mapData))
	for i := range visited {
		visited[i] = make([]bool, len(mapData[0]))
	}
	result := 0
	for y := 0; y < len(mapData); y++ {
		for x := 0; x < len(mapData[0]); x++ {
			if region := measureRegion(mapData, x, y, visited); region != nil {
				result += region[area] * region[sides]
			}
		}
	}
	return strconv.Itoa(result), nil
}

func parseMap(input string) [][]rune {
	lines := make([][]rune, 0)
	start := 0
	for i, c := range input {
		if c == '\n' {
			if i > start {
				lines = append(lines, []rune(input[start:i]))
			}
			start = i + 1
		}
	}
	if start < len(input) {
		lines = append(lines, []rune(input[start:]))
	}
	return lines
}

func measureRegion(mapData [][]rune, x, y int, visited [][]bool) []int {
	if visited[y][x] {
		return nil
	}
	region := make([]int, 3)
	measureRegionRec(mapData, x, y, visited, mapData[y][x], region)
	return region
}

func measureRegionRec(mapData [][]rune, x, y int, visited [][]bool, plant rune, region []int) {
	if visited[y][x] {
		return
	}
	region[area]++
	visited[y][x] = true
	width := len(mapData[0])
	height := len(mapData)
	// up
	if y > 0 && mapData[y-1][x] == plant {
		measureRegionRec(mapData, x, y-1, visited, plant, region)
	} else {
		region[perimeter]++
		sideContinuation := x > 0 && mapData[y][x-1] == plant && (y == 0 || mapData[y-1][x-1] != plant)
		if !sideContinuation {
			region[sides]++
		}
	}
	// right
	if x < width-1 && mapData[y][x+1] == plant {
		measureRegionRec(mapData, x+1, y, visited, plant, region)
	} else {
		region[perimeter]++
		sideContinuation := y > 0 && mapData[y-1][x] == plant && (x == width-1 || mapData[y-1][x+1] != plant)
		if !sideContinuation {
			region[sides]++
		}
	}
	// down
	if y < height-1 && mapData[y+1][x] == plant {
		measureRegionRec(mapData, x, y+1, visited, plant, region)
	} else {
		region[perimeter]++
		sideContinuation := x < width-1 && mapData[y][x+1] == plant && (y == height-1 || mapData[y+1][x+1] != plant)
		if !sideContinuation {
			region[sides]++
		}
	}
	// left
	if x > 0 && mapData[y][x-1] == plant {
		measureRegionRec(mapData, x-1, y, visited, plant, region)
	} else {
		region[perimeter]++
		sideContinuation := y < height-1 && mapData[y+1][x] == plant && (x == 0 || mapData[y+1][x-1] != plant)
		if !sideContinuation {
			region[sides]++
		}
	}
}
