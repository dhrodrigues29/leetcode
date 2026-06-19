package main

import (
	"fmt"
	"reflect"
)

func longestPalindrome(s string) string {

	return s
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
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "asdfghjkllkjhgfdsa",
			expected: "asdfghjkl",
		},
		{
			input:    "olskdmbqwertyytrewqurijam",
			expected: "qwerty",
		},
	}

	for i, tc := range tests {

		result := longestPalindrome(tc.input)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Nums1:    %v\n", tc.input)
			fmt.Printf("Expected: %.5f\n", tc.expected)
			fmt.Printf("Got:      %.5f\n", result)
		}
	}
}

func main() {
	runTests()
}
