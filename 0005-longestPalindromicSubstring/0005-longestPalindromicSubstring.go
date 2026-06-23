package main

import (
	"fmt"
	"reflect"
)

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	start := 0
	maxLen := 1

	expand := func(left, right int) {

		for left >= 0 &&
			right < len(s) &&
			s[left] == s[right] {

			currentLen := right - left + 1

			if currentLen > maxLen {
				maxLen = currentLen
				start = left
			}

			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {

		// odd center
		expand(i, i)

		// even center
		expand(i, i+1)
	}

	return s[start : start+maxLen]
}

type testCase struct {
	input    string
	expected string
}

func runTests() {

	tests := []testCase{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "aaaaaaababad",
			expected: "aaaaaaa",
		},
		{
			input:    "cbbbd",
			expected: "bbb",
		},
		{
			input:    "asdfghjkllkjhgfdsa",
			expected: "asdfghjkllkjhgfdsa",
		},
		{
			input:    "olskdmbqwertyytrwequrijam",
			expected: "rtyytr",
		},
	}

	for i, tc := range tests {

		result := longestPalindrome(tc.input)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Nums1:    %v\n", tc.input)
			fmt.Printf("Expected: %v\n", tc.expected)
			fmt.Printf("Got:      %v\n", result)
		}
	}
}

func main() {
	runTests()
}
