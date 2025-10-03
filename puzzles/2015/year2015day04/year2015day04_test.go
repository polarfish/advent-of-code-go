package year2015day04

import "testing"

func TestYear2015day04Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "282749", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015day04Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "9962624", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
