/*
Title: Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero,
except the number 0 itself.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import (
	"fmt"
	"reflect"
)

// my solution using recursion
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return recursiveSum(l1, l2, 0)
}

func recursiveSum(l1 *ListNode, l2 *ListNode, carry int) *ListNode {

	sum := carry

	if l1 != nil {
		sum += l1.Val
	}

	if l2 != nil {
		sum += l2.Val
	}

	if l1 == nil && l2 == nil && sum == 0 {
		return nil
	}

	store := sum % 10     // stores the unit
	nextCarry := sum / 10 //stores the tens

	node := &ListNode{
		Val: store,
	}

	var next1 *ListNode
	var next2 *ListNode

	if l1 != nil {
		next1 = l1.Next
	}

	if l2 != nil {
		next2 = l2.Next
	}

	node.Next = recursiveSum(next1, next2, nextCarry)

	return node
}

// helper list builder
func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for _, v := range values[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	return head
}

func listToSlice(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

// 2. Test structure
type testCase struct {
	l1       []int
	l2       []int
	expected []int
}

// 3. Runner
func runTests() {

	tests := []testCase{
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for i, tc := range tests {

		l1 := buildList(tc.l1)
		l2 := buildList(tc.l2)

		result := addTwoNumbers(l1, l2)

		got := listToSlice(result)

		if reflect.DeepEqual(got, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Expected: %v\n", tc.expected)
			fmt.Printf("Got:      %v\n", got)
		}
	}
}

func main() {
	runTests()
}
