package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello  world   ", 
			expected: []string{"hello", "world"}, 
		},
		{
			input: "   how are   YOU", 
			expected: []string{"how", "are", "you"}, 
		},
		{
			input: "   no WAY it s YOU   ", 
			expected: []string{"no", "way", "it", "s", "you"}, 
		},
	}
	for _, c := range cases{
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("got length %d but expected length %d for input %q, expected: %v  got: %v",
				len(actual), len(c.expected), c.input, c.expected, actual)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("mistmatch in words actual: %s and expected: %s", word, expectedWord)
			}
		}
	}
}
