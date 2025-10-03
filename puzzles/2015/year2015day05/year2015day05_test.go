package year2015day05

import "testing"

func TestYear2015day05Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "255", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015day05Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "55", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
