package main

import (
	"fmt"
	"reflect"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 1
	start := 0

	for start < len(s) {
		seen := make(map[byte]int)

		currentLen := 0
		nextStart := len(s)

		for i := start; i < len(s); i++ {
			char := s[i]

			if previousIndex, exists := seen[char]; exists {
				if currentLen > maxLen {
					maxLen = currentLen
				}

				nextStart = previousIndex + 1
				break
			}

			seen[char] = i
			currentLen++
		}

		if currentLen > maxLen {
			maxLen = currentLen
		}

		if nextStart == len(s) {
			break
		}

		start = nextStart
	}

	return maxLen
}

type testCase struct {
	input    string
	expected int
}

func runTests() {

	tests := []testCase{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "abcdefgcnmkl",
			expected: 7,
		},
		{
			input:    "abba",
			expected: 2,
		},
		{
			input:    "abcdbcabdbcbasbdbcbabsbdbcbasbdasbcbasbdassabbcbasdbascbcbsdabcbaskhmgnfjetyascbabsvcbavcbasvqwertyuioplkjhmnbgcvdfsahgjnabshgjsbdnasbgnbc",
			expected: 24,
		},
	}

	for i, tc := range tests {

		result := lengthOfLongestSubstring(tc.input)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input:    %q\n", tc.input)
			fmt.Printf("Expected: %d\n", tc.expected)
			fmt.Printf("Got:      %d\n", result)
		}
	}
}

func main() {
	runTests()
}
