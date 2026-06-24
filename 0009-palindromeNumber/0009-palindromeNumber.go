package main

import (
	"fmt"
	"reflect"
)

func isPalindrome(x int) bool {

	// Your logic here

	return false
}

type testCase struct {
	input    int
	expected bool
}

func runTests() {

	tests := []testCase{
		{
			input:    121,
			expected: true,
		},
		{
			input:    -121,
			expected: false,
		},
		{
			input:    10,
			expected: false,
		},
		{
			input:    0,
			expected: true,
		},
		{
			input:    1,
			expected: true,
		},
		{
			input:    11,
			expected: true,
		},
		{
			input:    12321,
			expected: true,
		},
		{
			input:    123321,
			expected: true,
		},
		{
			input:    12345,
			expected: false,
		},
		{
			input:    1001,
			expected: true,
		},
		{
			input:    100,
			expected: false,
		},
		{
			input:    2147447412,
			expected: true,
		},
	}

	for i, tc := range tests {

		result := isPalindrome(tc.input)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input:    %v\n", tc.input)
			fmt.Printf("Expected: %v\n", tc.expected)
			fmt.Printf("Got:      %v\n", result)
		}
	}
}

func main() {
	runTests()
}
