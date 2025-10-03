package year2015day06

import (
	"testing"
)

func TestYear2015Day06Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "400410", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015Day06Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "15343601", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
