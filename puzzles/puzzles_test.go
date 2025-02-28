package puzzles

import "testing"

type testCase struct {
	input    string
	expected string
}

func runTests(t *testing.T, testFunc func(input string) string, tests map[string]testCase) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := testFunc(test.input)
			if actual != test.expected {
				t.Errorf("\nexpected:\n%s\nactual:\n%s", test.expected, actual)
			}
		})
	}
}
