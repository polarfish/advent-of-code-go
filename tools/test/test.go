package test

import (
	"testing"
)

func Assert(t *testing.T, want string, part func(string) (string, error), input string) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("\nwant:\n%s\ngot panic:\n%s", want, r)
		}
	}()
	got, err := part(input)
	if err != nil {
		t.Errorf("\nwant:\n%s\ngot error:\n%s", want, err.Error())
	} else if got != want {
		t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
	}
}
