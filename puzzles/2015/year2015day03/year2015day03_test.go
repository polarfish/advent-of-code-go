package year2015day03

import "testing"

func TestYear2015day03Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "2572", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015day03Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "2631", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
