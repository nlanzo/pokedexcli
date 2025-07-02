package main

import "testing"



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
			input:    "  abcd efg hijk  e",
			expected: []string{"abcd", "efg", "hijk", "e"},
		},
		{
			input:    "  pikAchu gRimeR MeowTh  ",
			expected: []string{"pikachu", "grimer", "meowth"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %v, but got %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Expected %v, but got %v", expectedWord, word)
			}
		}
	}
}