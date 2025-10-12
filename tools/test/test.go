package test

import "testing"

func Assert(t *testing.T, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("\nwant:\n%s\ngot:\n%s", want, got)
	}
}
