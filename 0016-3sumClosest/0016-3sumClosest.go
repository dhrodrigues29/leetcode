package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	if len(nums) < 3 {
		return 0
	}

	//create return variable
	closest := nums[0] + nums[1] + nums[2]

	//sort the array
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		left := i + 1
		right := len(nums) - 1

		for left < right {

			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return sum
			}

			if sum < target {
				left++

			} else if sum > target {
				right--
			}

			currentDiff := sum - target
			if currentDiff < 0 {
				currentDiff = -currentDiff
			}

			closestDiff := closest - target
			if closestDiff < 0 {
				closestDiff = -closestDiff
			}

			if currentDiff < closestDiff {
				closest = sum
			}

		}

	}

	return closest

}

type testCase struct {
	input    []int
	target   int
	expected int
}

func runTests() {

	tests := []testCase{
		{
			input:    []int{-1, 2, 1, -4},
			target:   1,
			expected: 2,
		},
		{
			input:    []int{0, 0, 0},
			target:   1,
			expected: 0,
		},
		{
			input:    []int{1, 1, 1},
			target:   0,
			expected: 3,
		},
		{
			input:    []int{-1, -1, -1},
			target:   -5,
			expected: -3,
		},
		{
			input:    []int{-1, 0, 1},
			target:   0,
			expected: 0,
		},
		{
			input:    []int{-3, -2, -5, 3, -4},
			target:   -1,
			expected: -2,
		},
		{
			input:    []int{1, 2, 5, 10, 11},
			target:   12,
			expected: 13,
		},
		{
			input:    []int{-7, -5, -3, -1},
			target:   -10,
			expected: -11,
		},
		{
			input:    []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target:   -2,
			expected: -2,
		},
		{
			input:    []int{0, 2, 1, -3},
			target:   1,
			expected: 0,
		},
		{
			input:    []int{1, 6, 9, 14, 16, 70},
			target:   81,
			expected: 80,
		},
		{
			input:    []int{-1000, -999, 999, 1000, 1},
			target:   0,
			expected: 0,
		},
		{
			input:    []int{-8, -6, -5, -3, -1},
			target:   -13,
			expected: -14,
		},
		{
			input:    []int{5, 2, 7, 1, 8},
			target:   15,
			expected: 15,
		},
		{
			input:    []int{-2, 0, 1, 2},
			target:   2,
			expected: 1,
		},
		{
			input:    []int{-4, -1, 1, 2},
			target:   1,
			expected: 2,
		},
		{
			input:    []int{3, 0, -2, -1, 1, 2},
			target:   4,
			expected: 4,
		},
		{
			input:    []int{1000, 1000, 1000},
			target:   -10000,
			expected: 3000,
		},
		{
			input:    []int{-1000, -1000, -1000},
			target:   10000,
			expected: -3000,
		},
		{
			input:    []int{-10, -5, -2, 0, 4, 8, 12},
			target:   7,
			expected: 7,
		},
	}

	var failed []int

	for i, tc := range tests {

		result := threeSumClosest(tc.input, tc.target)

		if result == tc.expected {
			fmt.Printf("Test %2d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %2d FAILED\n", i+1)
			failed = append(failed, i)
		}
	}

	if len(failed) > 0 {

		fmt.Println("\nFailed Test Details")
		fmt.Println("-------------------")

		for _, i := range failed {

			tc := tests[i]
			result := threeSumClosest(tc.input, tc.target)

			fmt.Printf("Test %d\n", i+1)
			fmt.Printf("Input:    %v\n", tc.input)
			fmt.Printf("Target:   %d\n", tc.target)
			fmt.Printf("Expected: %d\n", tc.expected)
			fmt.Printf("Got:      %d\n\n", result)
		}
	}

	fmt.Printf("Summary: %d/%d tests passed\n",
		len(tests)-len(failed), len(tests))
}

func main() {
	runTests()
}
