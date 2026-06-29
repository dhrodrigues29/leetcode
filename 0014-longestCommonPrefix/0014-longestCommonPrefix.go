package main

import (
	"fmt"
	"reflect"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	currentLongestPrefix := strs[0]

	for i := 0; i < len(strs)-1; i++ {

		currentString := currentLongestPrefix
		nextString := strs[i+1]

		newPrefix := ""
		limit := min(len(currentString), len(nextString))

		for j := 0; j < limit; j++ {

			if currentString[j] != nextString[j] {
				break
			}

			newPrefix += string(currentString[j])
		}

		currentLongestPrefix = newPrefix

		if currentLongestPrefix == "" {
			return ""
		}

	}

	return currentLongestPrefix
}

type testCase struct {
	input    []string
	expected string
}

func runTests() {

	tests := []testCase{
		{
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			input:    []string{"interspecies", "interstellar", "interstate"},
			expected: "inters",
		},
		{
			input:    []string{"throne", "throne"},
			expected: "throne",
		},
		{
			input:    []string{"", ""},
			expected: "",
		},
		{
			input:    []string{""},
			expected: "",
		},
		{
			input:    []string{"a"},
			expected: "a",
		},
		{
			input:    []string{"ab", "a"},
			expected: "a",
		},
		{
			input:    []string{"a", "ab"},
			expected: "a",
		},
		{
			input:    []string{"abc", "abc", "abc"},
			expected: "abc",
		},
		{
			input:    []string{"prefix", "pre", "prevent"},
			expected: "pre",
		},
		{
			input:    []string{"apple", "banana", "cherry"},
			expected: "",
		},
		{
			input:    []string{"test"},
			expected: "test",
		},
		{
			input:    []string{},
			expected: "",
		},
		{
			input:    []string{"aaa", "aa", "aaa"},
			expected: "aa",
		},
	}

	for i, tc := range tests {

		result := longestCommonPrefix(tc.input)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input:    %v\n", tc.input)
			fmt.Printf("Expected: %q\n", tc.expected)
			fmt.Printf("Got:      %q\n", result)
		}
	}
}

func main() {
	runTests()
}
