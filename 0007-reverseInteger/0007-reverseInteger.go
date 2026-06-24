package main

import (
	"fmt"
	"reflect"
)

func reverse(x int) int {

	// TODO: Implement solution

	return 0
}

type testCase struct {
	input    int
	expected int
}

func runTests() {

	tests := []testCase{
		{
			input:    123,
			expected: 321,
		},
		{
			input:    -123,
			expected: -321,
		},
		{
			input:    120,
			expected: 21,
		},
		{
			input:    0,
			expected: 0,
		},
		{
			input:    5,
			expected: 5,
		},
		{
			input:    -5,
			expected: -5,
		},
		{
			input:    1000,
			expected: 1,
		},
		{
			input:    -1000,
			expected: -1,
		},
		{
			input:    1534236469, // overflow when reversed
			expected: 0,
		},
		{
			input:    -1534236469, // overflow when reversed
			expected: 0,
		},
		{
			input:    1463847412, // valid reverse
			expected: 2147483641,
		},
		{
			input:    -1463847412, // valid reverse
			expected: -2147483641,
		},
	}

	for i, tc := range tests {

		result := reverse(tc.input)

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
