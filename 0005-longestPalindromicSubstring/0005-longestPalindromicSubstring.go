package main

import (
	"fmt"
	"reflect"
)

func longestPalindrome(s string) int {

	maxLen := 1
	currentMaxLen := 1

	for i := 0; i < len(s); i++ {
		if i == 0 && i+1 != len(s) {
			for string(s[i]) == string(s[i+1]) {
				currentMaxLen += 1
				i++
				if i >= len(s) {
					break
				}
			}
			maxLen = currentMaxLen
			currentMaxLen = 1
		}

		for j := 1; j < len(s); j++ {
			if !(i-j < 0 || i+j-1 >= len(s)) {
				if string(s[i-j]) == string(s[i+j-1]) {
					currentMaxLen += 2
					if currentMaxLen > maxLen {
						maxLen = currentMaxLen
					}
				}
			}
		}
	}
	return maxLen
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
