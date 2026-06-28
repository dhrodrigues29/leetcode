package main

import (
	"fmt"
	"reflect"
)

func longestCommonPrefix(strs []string) string {

	return ""
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