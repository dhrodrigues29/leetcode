package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func isPalindrome(x int) bool {

	// if x is negative, it cannot be palindrome
	if x < 0 {
		return false
	}

	// if x is above or bellow 32-bit range, it is out of the constrains range
	if x < -2147483648 || x > 2147483647 {
		return false
	}

	//convert int to string
	s := strconv.Itoa(x)

	// palindrome numbers must always have equal starter and ending numbers, so we track both ends of the number togheter and compare them
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {

		if (left-right) == 0 && s[left] == s[right] {
			return true

		}

		if (left-right) == -1 && s[left] == s[right] {
			return true

		}

		if s[left] != s[right] {
			return false
		}
	}

	//if gets here, then the number is a palindrome since all return false condition were never reached
	return true
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
