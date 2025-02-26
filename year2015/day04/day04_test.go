package year2015day04

import "testing"

func TestPart1(t *testing.T) {
	got := Part1(input)
	t.Log("Part 1:", got)
	want := "282749"
	if got != want {
		t.Errorf("Part1: \ngot %s\nwant %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(input)
	t.Log("Part 2:", got)
	want := "9962624"
	if got != want {
		t.Errorf("Part2: \ngot %s\nwant %s", got, want)
	}
}
