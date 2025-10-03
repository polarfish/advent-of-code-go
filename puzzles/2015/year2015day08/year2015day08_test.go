package year2015day08

import (
	"testing"
)

func TestYear2015Day08Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "0", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015Day08Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "0", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
