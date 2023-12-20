package utils

import (
	"testing"
)

func TestGetCmdFromPrompt(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "Hello wOrld", expected: []string{"hello", "world"}},
		{input: "Lorem ipSum", expected: []string{"lorem", "ipsum"}},
	}

	for _, cs := range cases {
		slicedPrompt, _ := GetCmdFromPrompt(cs.input)
		for i := range slicedPrompt {
			actual := slicedPrompt[i]
			expected := cs.expected[i]
			if actual != expected {
				t.Errorf("%s doesn't equals %s", actual, expected)
			}
		}
	}
}
