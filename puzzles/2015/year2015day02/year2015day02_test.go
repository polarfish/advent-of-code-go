package year2015day02

import "testing"

func TestYear2015day02Part1(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "1606483", part1(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}

func TestYear2015day02Part2(t *testing.T) {
	t.Run("input", func(t *testing.T) {
		if want, got := "3842356", part2(input); got != want {
			t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
		}
	})
}
