package year2015day01

import "testing"

func TestYear2015day01Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "280", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015day01Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "1797", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
