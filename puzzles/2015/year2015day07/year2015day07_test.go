package year2015day07

import (
	"testing"
)

func TestYear2015Day07Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "3176", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015Day07Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "14710", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
