package main

import (
	"fmt"
	"reflect"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	maxNewArraySize := (len(nums1) + len(nums2))

	neededNewArraySize := maxNewArraySize / 2

	newArray := make([]int, maxNewArraySize)

	i, j, k := 0, 0, 0

	for i < len(nums1) || j < len(nums2) {

		if j >= len(nums2) || (i < len(nums1) && nums1[i] < nums2[j]) {
			newArray[k] = nums1[i]
			i++
		} else {
			newArray[k] = nums2[j]
			j++
		}

		k++

		if k == neededNewArraySize+1 {

			if maxNewArraySize%2 == 1 {
				return float64(newArray[k-1])
			} else {
				return float64(newArray[k-1]+newArray[k-2]) / 2.0
			}
		}
	}

	return 0
}

type testCase struct {
	nums1    []int
	nums2    []int
	expected float64
}

func runTests() {

	tests := []testCase{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3},
			expected: 2.0,
		},
		{
			nums1:    []int{3},
			nums2:    []int{1, 2, 4},
			expected: 2.5,
		},
		{
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
		{
			nums1:    []int{2},
			nums2:    []int{1},
			expected: 1.5,
		},
		{
			nums1:    []int{-5, -3, -1},
			nums2:    []int{-2, 0, 2},
			expected: -1.5,
		},
	}

	for i, tc := range tests {

		result := findMedianSortedArrays(tc.nums1, tc.nums2)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Nums1:    %v\n", tc.nums1)
			fmt.Printf("Nums2:    %v\n", tc.nums2)
			fmt.Printf("Expected: %.5f\n", tc.expected)
			fmt.Printf("Got:      %.5f\n", result)
		}
	}
}

func main() {
	runTests()
}
