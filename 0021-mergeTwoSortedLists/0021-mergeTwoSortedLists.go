package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

}

type testCase struct {
	list1    []int
	list2    []int
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
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
		{
			list1:    []int{0},
			list2:    []int{},
			expected: []int{0},
		},
		{
			list1:    []int{1},
			list2:    []int{2},
			expected: []int{1, 2},
		},
		{
			list1:    []int{2},
			list2:    []int{1},
			expected: []int{1, 2},
		},
		{
			list1:    []int{1, 1, 1},
			list2:    []int{1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			list1:    []int{-10, -5, 0},
			list2:    []int{-8, -3, 2},
			expected: []int{-10, -8, -5, -3, 0, 2},
		},
		{
			list1:    []int{-5, -3, -1},
			list2:    []int{0, 2, 4},
			expected: []int{-5, -3, -1, 0, 2, 4},
		},
		{
			list1:    []int{0, 2, 4},
			list2:    []int{-5, -3, -1},
			expected: []int{-5, -3, -1, 0, 2, 4},
		},
		{
			list1:    []int{1, 3, 5},
			list2:    []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			list1:    []int{2, 4, 6},
			list2:    []int{1, 3, 5},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			list1:    []int{1, 2, 3},
			list2:    []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			list1:    []int{4, 5, 6},
			list2:    []int{1, 2, 3},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			list1:    []int{1, 2, 2, 4},
			list2:    []int{2, 2, 3},
			expected: []int{1, 2, 2, 2, 2, 3, 4},
		},
		{
			list1:    []int{-100},
			list2:    []int{100},
			expected: []int{-100, 100},
		},
		{
			list1:    []int{100},
			list2:    []int{-100},
			expected: []int{-100, 100},
		},
		{
			list1:    []int{-3, -2, -1},
			list2:    []int{-6, -5, -4},
			expected: []int{-6, -5, -4, -3, -2, -1},
		},
		{
			list1:    []int{5},
			list2:    []int{1, 2, 3, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			list1:    []int{1, 2, 3, 4, 6},
			list2:    []int{5},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	var failed []int

	for i, tc := range tests {

		result := listToSlice(
			mergeTwoLists(
				buildList(tc.list1),
				buildList(tc.list2),
			),
		)

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

			result := listToSlice(
				mergeTwoLists(
					buildList(tc.list1),
					buildList(tc.list2),
				),
			)

			fmt.Printf("Test %d\n", i+1)
			fmt.Printf("List1:    %v\n", tc.list1)
			fmt.Printf("List2:    %v\n", tc.list2)
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
