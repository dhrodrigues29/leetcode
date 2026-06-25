package main

import (
	"fmt"
	"reflect"
)

func isMatch(s string, p string) bool {

	// TODO: Implement solution

	return false
}

type testCase struct {
	s        string
	p        string
	expected bool
}

func runTests() {

	tests := []testCase{
		{
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			s:        "mississippi",
			p:        "mis*is*p*.",
			expected: false,
		},
		{
			s:        "abc",
			p:        "abc",
			expected: true,
		},
		{
			s:        "abc",
			p:        "abd",
			expected: false,
		},
		{
			s:        "aaa",
			p:        "a*a",
			expected: true,
		},
		{
			s:        "aaa",
			p:        "ab*a*c*a",
			expected: true,
		},
		{
			s:        "abcd",
			p:        "d*",
			expected: false,
		},
		{
			s:        "a",
			p:        ".",
			expected: true,
		},
		{
			s:        "ab",
			p:        "..",
			expected: true,
		},
		{
			s:        "ab",
			p:        "...",
			expected: false,
		},
		{
			s:        "aaaaaaaaaaaaab",
			p:        "a*a*a*a*a*a*a*a*a*a*c",
			expected: false,
		},
	}

	for i, tc := range tests {

		result := isMatch(tc.s, tc.p)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input s:  %q\n", tc.s)
			fmt.Printf("Input p:  %q\n", tc.p)
			fmt.Printf("Expected: %t\n", tc.expected)
			fmt.Printf("Got:      %t\n", result)
		}
	}
}

func main() {
	runTests()
}
