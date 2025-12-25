package year2025day09

import (
	"cmp"
	_ "embed"
	slices "slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day09.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/9
	registry.AddSolution(2025, 9, "Movie Theater", input, part1, part2)
}

func part1(input string) (string, error) {
	var result int64

	lines := utils.Lines(input)
	points, err := parseInput(lines)
	if err != nil {
		return "", utils.ErrBadInput
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			w, h := abs(p1.x-p2.x)+1, abs(p1.y-p2.y)+1
			area := int64(w) * int64(h)
			result = max(result, area)
		}
	}

	return strconv.FormatInt(result, 10), nil
}

func part2(input string) (string, error) {
	var result int64

	lines := utils.Lines(input)
	points, err := parseInput(lines)
	if err != nil {
		return "", utils.ErrBadInput
	}

	var dir int
	if points[0].x == points[len(points)-1].x {
		if points[0].y < points[len(points)-1].y {
			dir = top
		} else {
			dir = bottom
		}
	} else {
		if points[0].x < points[len(points)-1].x {
			dir = right
		} else {
			dir = left
		}
	}
	corners := [][]point{{}, {}, {}, {}}
	verticalEdges := make([]edge, 0)
	horizontalEdges := make([]edge, 0)
	for i := 0; i < len(points); i++ {
		p1, p2 := points[i], points[(i+1)%len(points)]
		switch dir {
		case top:
			if p2.x > p1.x {
				dir = right
				corners[topLeft] = append(corners[topLeft], p1)
				horizontalEdges = append(horizontalEdges, edge{p1.x, p1.y, p2.x, p2.y})
			} else {
				dir = left
				corners[topLeft] = append(corners[topLeft], p1)
				corners[bottomLeft] = append(corners[bottomLeft], p1)
				corners[bottomRight] = append(corners[bottomRight], p1)
				horizontalEdges = append(horizontalEdges, edge{p2.x, p2.y, p1.x, p1.y})
			}
		case bottom:
			if p2.x > p1.x {
				dir = right
				corners[topLeft] = append(corners[topLeft], p1)
				corners[topRight] = append(corners[topRight], p1)
				corners[bottomRight] = append(corners[bottomRight], p1)
				horizontalEdges = append(horizontalEdges, edge{p1.x, p1.y, p2.x, p2.y})
			} else {
				dir = left
				corners[bottomRight] = append(corners[bottomRight], p1)
				horizontalEdges = append(horizontalEdges, edge{p2.x, p2.y, p1.x, p1.y})
			}
		case right:
			if p2.y > p1.y {
				dir = bottom
				corners[topRight] = append(corners[topRight], p1)
				verticalEdges = append(verticalEdges, edge{p1.x, p1.y, p2.x, p2.y})
			} else {
				dir = top
				corners[topRight] = append(corners[topRight], p1)
				corners[topLeft] = append(corners[topLeft], p1)
				corners[bottomLeft] = append(corners[bottomLeft], p1)
				verticalEdges = append(verticalEdges, edge{p2.x, p2.y, p1.x, p1.y})
			}
		case left:
			if p2.y > p1.y {
				dir = bottom
				corners[topRight] = append(corners[topRight], p1)
				corners[bottomRight] = append(corners[bottomRight], p1)
				corners[bottomRight] = append(corners[bottomRight], p1)
				verticalEdges = append(verticalEdges, edge{p1.x, p1.y, p2.x, p2.y})
			} else {
				dir = top
				corners[bottomLeft] = append(corners[bottomLeft], p1)
				verticalEdges = append(verticalEdges, edge{p2.x, p2.y, p1.x, p1.y})
			}
		}
	}

	slices.SortFunc(verticalEdges, func(a, b edge) int {
		return cmp.Compare(a.x1, b.x1)
	})

	slices.SortFunc(horizontalEdges, func(a, b edge) int {
		return cmp.Compare(a.y1, b.y1)
	})

	for _, tl := range corners[topLeft] {
		for _, br := range corners[bottomRight] {
			if tl.x < br.x && tl.y < br.y {
				w, h := br.x-tl.x+1, br.y-tl.y+1
				area := int64(w) * int64(h)
				if area > result && isValidRectangle(verticalEdges, horizontalEdges, tl, br) {
					result = area
				}
			}
		}
	}

	for _, bl := range corners[bottomLeft] {
		for _, tr := range corners[topRight] {
			if bl.x < tr.x && bl.y > tr.y {
				w, h := tr.x-bl.x+1, bl.y-tr.y+1
				area := int64(w) * int64(h)
				if area > result && isValidRectangle(verticalEdges, horizontalEdges, bl, tr) {
					result = area
				}
			}
		}
	}

	return strconv.FormatInt(result, 10), nil
}

func isValidRectangle(verticalEdges []edge, horizontalEdges []edge, p1 point, p2 point) bool {
	rxl := min(p1.x, p2.x)
	rxr := max(p1.x, p2.x)
	ryt := min(p1.y, p2.y)
	ryb := max(p1.y, p2.y)

	for _, e := range verticalEdges {
		if rxl < e.x2 && rxr > e.x1 && ryt < e.y2 && ryb > e.y1 {
			return false
		}
	}

	for _, e := range horizontalEdges {
		if rxl < e.x2 && rxr > e.x1 && ryt < e.y2 && ryb > e.y1 {
			return false
		}
	}

	return true
}

const topLeft, topRight, bottomRight, bottomLeft = 0, 1, 2, 3
const top, right, bottom, left = 0, 1, 2, 3

type point struct {
	x, y int
}

type edge struct {
	x1, y1, x2, y2 int
}

func parseInput(lines []string) ([]point, error) {
	points := make([]point, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		points[i] = point{x, y}
	}
	return points, nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
