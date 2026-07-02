package main

import (
	"fmt"
)

func isValid(s string) bool {

	if len(s)%2 == 1 {
		return false
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {

		switch s[i] {

		case '(', '[', '{':

			stack = append(stack, s[i])

		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]

		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]

		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0

}

type testCase struct {
	input    string
	expected bool
}

func runTests() {

	tests := []testCase{
		{
			input:    "()",
			expected: true,
		},
		{
			input:    "()[]{}",
			expected: true,
		},
		{
			input:    "(]",
			expected: false,
		},
		{
			input:    "([])",
			expected: true,
		},
		{
			input:    "([)]",
			expected: false,
		},
		{
			input:    "{[]}",
			expected: true,
		},
		{
			input:    "(",
			expected: false,
		},
		{
			input:    ")",
			expected: false,
		},
		{
			input:    "(((",
			expected: false,
		},
		{
			input:    ")))",
			expected: false,
		},
		{
			input:    "((()))",
			expected: true,
		},
		{
			input:    "(()",
			expected: false,
		},
		{
			input:    "())",
			expected: false,
		},
		{
			input:    "[({})]",
			expected: true,
		},
		{
			input:    "[({})",
			expected: false,
		},
		{
			input:    "}{",
			expected: false,
		},
		{
			input:    "[]{}()",
			expected: true,
		},
		{
			input:    "[[[[]]]]",
			expected: true,
		},
		{
			input:    "[(])",
			expected: false,
		},
		{
			input:    "{{[[(())]]}}",
			expected: true,
		},
	}

	var failed []int

	for i, tc := range tests {

		result := isValid(tc.input)

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
			result := isValid(tc.input)

			fmt.Printf("Test %d\n", i+1)
			fmt.Printf("Input:    %q\n", tc.input)
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
