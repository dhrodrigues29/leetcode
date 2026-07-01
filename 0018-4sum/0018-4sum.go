package main

import (
	"fmt"
	"reflect"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

}

type testCase struct {
	input    []int
	target   int
	expected [][]int
}

func normalize(result [][]int) [][]int {

	for i := range result {
		sort.Ints(result[i])
	}

	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < 4; k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
	})

	return result
}

func runTests() {

	tests := []testCase{
		{
			input:    []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			input:    []int{2, 2, 2, 2, 2},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}},
		},
		{
			input:    []int{},
			target:   0,
			expected: [][]int{},
		},
		{
			input:    []int{1},
			target:   1,
			expected: [][]int{},
		},
		{
			input:    []int{1, 2, 3},
			target:   6,
			expected: [][]int{},
		},
		{
			input:    []int{0, 0, 0, 0},
			target:   0,
			expected: [][]int{{0, 0, 0, 0}},
		},
		{
			input:    []int{0, 0, 0, 0, 0},
			target:   0,
			expected: [][]int{{0, 0, 0, 0}},
		},
		{
			input:    []int{-3, -1, 0, 2, 4, 5},
			target:   2,
			expected: [][]int{{-3, -1, 2, 4}},
		},
		{
			input:    []int{-2, -1, 0, 0, 1, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			input:    []int{-2, -2, -2, 2, 2, 2},
			target:   0,
			expected: [][]int{{-2, -2, 2, 2}},
		},
		{
			input:    []int{-5, -4, -3, -2, -1},
			target:   -10,
			expected: [][]int{{-4, -3, -2, -1}},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			target:   50,
			expected: [][]int{},
		},
		{
			input:    []int{1000000000, 1000000000, -1000000000, -1000000000, 0},
			target:   0,
			expected: [][]int{{-1000000000, -1000000000, 1000000000, 1000000000}},
		},
		{
			input:    []int{-1, -1, -1, -1, 2, 2, 2, 2},
			target:   2,
			expected: [][]int{{-1, -1, 2, 2}},
		},
		{
			input:    []int{-4, -1, -1, 0, 1, 2},
			target:   -4,
			expected: [][]int{{-4, -1, -1, 2}},
		},
		{
			input:    []int{5, 5, 3, 5, 1, -5, 1, -2},
			target:   4,
			expected: [][]int{{-5, 1, 3, 5}},
		},
		{
			input:  []int{-3, -2, -1, 0, 0, 1, 2, 3},
			target: 0,
			expected: [][]int{
				{-3, -2, 2, 3},
				{-3, -1, 1, 3},
				{-3, 0, 0, 3},
				{-3, 0, 1, 2},
				{-2, -1, 0, 3},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			input:    []int{1, 1, 1, 1, 1, 1},
			target:   4,
			expected: [][]int{{1, 1, 1, 1}},
		},
		{
			input:    []int{-10, -5, 0, 5, 10},
			target:   0,
			expected: [][]int{{-10, -5, 5, 10}},
		},
		{
			input:  []int{-8, -6, -4, -2, 2, 4, 6, 8},
			target: 0,
			expected: [][]int{
				{-8, -6, 6, 8},
				{-8, -4, 4, 8},
				{-8, -2, 2, 8},
				{-8, -2, 4, 6},
				{-6, -4, 2, 8},
				{-6, -4, 4, 6},
				{-6, -2, 2, 6},
				{-4, -2, 2, 4},
			},
		},
	}

	var failed []int

	for i, tc := range tests {

		result := normalize(fourSum(tc.input, tc.target))
		expected := normalize(tc.expected)

		if reflect.DeepEqual(result, expected) {
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
			result := normalize(fourSum(tc.input, tc.target))
			expected := normalize(tc.expected)

			fmt.Printf("Test %d\n", i+1)
			fmt.Printf("Input:    %v\n", tc.input)
			fmt.Printf("Target:   %d\n", tc.target)
			fmt.Printf("Expected: %v\n", expected)
			fmt.Printf("Got:      %v\n\n", result)
		}
	}

	fmt.Printf("Summary: %d/%d tests passed\n",
		len(tests)-len(failed), len(tests))
}

func main() {
	runTests()
}
