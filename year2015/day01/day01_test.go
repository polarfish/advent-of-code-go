package year2015day01

import "testing"

func TestPart1(t *testing.T) {
	want := "280"
	got := Part1(&input)
	if got != want {
		t.Errorf("Part1: \ngot %s\nwant %s", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "1797"
	got := Part2(&input)
	if got != want {
		t.Errorf("Part2: \ngot %s\nwant %s", got, want)
	}
}
