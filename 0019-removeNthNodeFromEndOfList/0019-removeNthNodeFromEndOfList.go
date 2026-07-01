package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	return nil
}

type testCase struct {
	input    []int
	n        int
	expected []int
}

func buildList(values []int) *ListNode {

	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

func listToSlice(head *ListNode) []int {

	result := []int{}

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func runTests() {

	tests := []testCase{
		{
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			input:    []int{1},
			n:        1,
			expected: []int{},
		},
		{
			input:    []int{1, 2},
			n:        1,
			expected: []int{1},
		},
		{
			input:    []int{1, 2},
			n:        2,
			expected: []int{2},
		},
		{
			input:    []int{1, 2, 3},
			n:        1,
			expected: []int{1, 2},
		},
		{
			input:    []int{1, 2, 3},
			n:        2,
			expected: []int{1, 3},
		},
		{
			input:    []int{1, 2, 3},
			n:        3,
			expected: []int{2, 3},
		},
		{
			input:    []int{5, 6, 7, 8},
			n:        4,
			expected: []int{6, 7, 8},
		},
		{
			input:    []int{5, 6, 7, 8},
			n:        3,
			expected: []int{5, 7, 8},
		},
		{
			input:    []int{5, 6, 7, 8},
			n:        2,
			expected: []int{5, 6, 8},
		},
		{
			input:    []int{5, 6, 7, 8},
			n:        1,
			expected: []int{5, 6, 7},
		},
		{
			input:    []int{0, 0, 0},
			n:        2,
			expected: []int{0, 0},
		},
		{
			input:    []int{1, 1, 1, 1},
			n:        3,
			expected: []int{1, 1, 1},
		},
		{
			input:    []int{10, 20, 30, 40, 50, 60},
			n:        6,
			expected: []int{20, 30, 40, 50, 60},
		},
		{
			input:    []int{10, 20, 30, 40, 50, 60},
			n:        5,
			expected: []int{10, 30, 40, 50, 60},
		},
		{
			input:    []int{10, 20, 30, 40, 50, 60},
			n:        4,
			expected: []int{10, 20, 40, 50, 60},
		},
		{
			input:    []int{10, 20, 30, 40, 50, 60},
			n:        3,
			expected: []int{10, 20, 30, 50, 60},
		},
		{
			input:    []int{10, 20, 30, 40, 50, 60},
			n:        2,
			expected: []int{10, 20, 30, 40, 60},
		},
		{
			input:    []int{10, 20, 30, 40, 50, 60},
			n:        1,
			expected: []int{10, 20, 30, 40, 50},
		},
		{
			input:    []int{42, 99},
			n:        2,
			expected: []int{99},
		},
	}

	var failed []int

	for i, tc := range tests {

		head := buildList(tc.input)
		result := listToSlice(removeNthFromEnd(head, tc.n))

		if reflect.DeepEqual(result, tc.expected) {
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
			head := buildList(tc.input)
			result := listToSlice(removeNthFromEnd(head, tc.n))

			fmt.Printf("Test %d\n", i+1)
			fmt.Printf("Input:    head=%v, n=%d\n", tc.input, tc.n)
			fmt.Printf("Expected: %v\n", tc.expected)
			fmt.Printf("Got:      %v\n\n", result)
		}
	}

	fmt.Printf("Summary: %d/%d tests passed\n",
		len(tests)-len(failed), len(tests))
}

func main() {
	runTests()
}
