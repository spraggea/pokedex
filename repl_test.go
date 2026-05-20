package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "Something is crazy!  ",
			expected: []string{"something", "is", "crazy!"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		if len(actual) != len(c.expected) {
			t.Errorf("Length Does Match %d , got %d ", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test

			if word != expectedWord {
				t.Errorf("Word does not match %s, got %s", expectedWord, word)
				//t.Error does not return anything, just outputds and fails test
				continue
			}
		}
	}
}
