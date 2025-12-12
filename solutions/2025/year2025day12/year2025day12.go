package year2025day12

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2025day12.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/12
	registry.AddSolution(2025, 12, "Christmas Tree Farm", input, part1, part2)
}

func part1(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	regionsParts := utils.Lines(parts[len(parts)-1])

	regions := make([]region, len(regionsParts))
	for i, regionPart := range regionsParts {
		regionSplit := strings.Split(regionPart, ": ")
		regionSizeSplit := strings.Split(regionSplit[0], "x")
		width, err := strconv.Atoi(regionSizeSplit[0])
		if err != nil {
			return "", utils.ErrBadInput
		}
		height, err := strconv.Atoi(regionSizeSplit[1])
		if err != nil {
			return "", utils.ErrBadInput
		}
		regionShapesSplit := strings.Split(regionSplit[1], " ")
		regionShapes := make([]int, len(regionShapesSplit))

		for j, shapeStr := range regionShapesSplit {
			shape, err := strconv.Atoi(shapeStr)
			if err != nil {
				return "", utils.ErrBadInput
			}
			regionShapes[j] = shape
		}

		regions[i] = region{
			width:  width,
			height: height,
			shapes: regionShapes,
		}
	}

	no := 0
	yes := 0
	maybe := 0
	for _, r := range regions {
		totalShapes := 0
		for _, s := range r.shapes {
			totalShapes += s
		}

		if r.width*r.height >= totalShapes*9 {
			yes++
		} else if r.width*r.height < totalShapes*7 {
			no++
		} else {
			maybe++
		}
	}

	fmt.Println(yes, no, maybe)

	return strconv.Itoa(yes), nil
}

func part2(input string) (string, error) {
	return strconv.Itoa(0), nil
}

type region struct {
	width, height int
	shapes        []int
}
