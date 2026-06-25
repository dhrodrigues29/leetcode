package main

import (
	"fmt"
	"reflect"
)

func maxArea(height []int) int {

	return 0

}

type testCase struct {
	height   []int
	expected int
}

func runTests() {

	tests := []testCase{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
		{
			height:   []int{1, 2, 3},
			expected: 2,
		},
		{
			height:   []int{5, 5, 5, 5},
			expected: 15,
		},
		{
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			height:   []int{1, 2, 1},
			expected: 2,
		},
		{
			height:   []int{2, 3, 4, 5, 18, 17, 6},
			expected: 17,
		},
		{
			height:   []int{1, 3, 2, 5, 25, 24, 5},
			expected: 24,
		},
	}

	for i, tc := range tests {

		fmt.Printf("Test %d STARTING\n", i+1)

		result := maxArea(tc.height)

		if reflect.DeepEqual(result, tc.expected) {

			fmt.Printf("Test %d PASSED\n", i+1)

		} else {

			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input:    %v\n", tc.height)
			fmt.Printf("Expected: %d\n", tc.expected)
			fmt.Printf("Got:      %d\n", result)
		}
	}
}

func main() {
	runTests()
}
