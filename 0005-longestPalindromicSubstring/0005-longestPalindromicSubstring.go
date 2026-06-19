package main

import (
	"fmt"
	"reflect"
)

func longestPalindrome(s string) string {

	maxLenSubstring := ""

	for i := 0; i < len(s); i++ {

		current := string(s[i])

		for j := 1; ; j++ {

			left := i - j
			right := i + j

			if left < 0 || right >= len(s) {
				break
			}

			if s[left] != s[right] {
				break
			}

			current = string(s[left]) + current + string(s[right])

			if len(current) > len(maxLenSubstring) {
				maxLenSubstring = current
			}
		}
	}

	return maxLenSubstring
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
			expected: "qwertyytrewq",
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
