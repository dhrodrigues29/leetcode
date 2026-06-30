package main

import (
	"fmt"
	"reflect"
	"sort"
)

func threeSum(nums []int) [][]int {

	//create return variable
	validTriplets := [][]int{}

	//sort the array
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {

			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++

			} else if sum > 0 {
				right--

			} else {
				validTriplets = append(validTriplets,
					[]int{nums[i], nums[left], nums[right]},
				)

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}

		}

	}

	return validTriplets

}

type testCase struct {
	input    []int
	expected [][]int
}

func normalize(result [][]int) [][]int {

	for i := range result {
		sort.Ints(result[i])
	}

	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < 3; k++ {
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
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			input:    []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			input:    []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			input:    []int{-2, 0, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			input:    []int{-2, -1, 3},
			expected: [][]int{{-2, -1, 3}},
		},
		{
			input:    []int{-1, -1, 2},
			expected: [][]int{{-1, -1, 2}},
		},
		{
			input:    []int{-1, -1, -1, 2, 2},
			expected: [][]int{{-1, -1, 2}},
		},
		{
			input:    []int{-4, -2, -2, 0, 1, 2, 2, 4},
			expected: [][]int{{-4, 0, 4}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
		{
			input:    []int{3, -2, 1, 0},
			expected: [][]int{},
		},
		{
			input:    []int{-5, 2, 3, 0},
			expected: [][]int{{-5, 2, 3}},
		},
		{
			input:    []int{-3, -2, -1},
			expected: [][]int{},
		},
		{
			input:    []int{1, 2, 3},
			expected: [][]int{},
		},
		{
			input:    []int{-100000, 0, 100000},
			expected: [][]int{{-100000, 0, 100000}},
		},
		{
			input:    []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			input:    []int{-1, 0, 1, 0},
			expected: [][]int{{-1, 0, 1}},
		},
		{
			input:    []int{-2, -2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			input:    []int{-7, -3, -2, 5, 10, -8},
			expected: [][]int{{-8, -2, 10}, {-7, -3, 10}, {-3, -2, 5}},
		},
		{
			input:    []int{-6, -1, 0, 1, 2, 4},
			expected: [][]int{{-6, 2, 4}, {-1, 0, 1}},
		},
		{
			input:    []int{-10, -7, -3, 0, 3, 7, 10},
			expected: [][]int{{-10, 0, 10}, {-10, 3, 7}, {-7, -3, 10}, {-7, 0, 7}, {-3, 0, 3}},
		},
	}

	var failed []int

	for i, tc := range tests {

		result := normalize(threeSum(tc.input))
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
			result := normalize(threeSum(tc.input))
			expected := normalize(tc.expected)

			fmt.Printf("Test %d\n", i+1)
			fmt.Printf("Input:    %v\n", tc.input)
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
