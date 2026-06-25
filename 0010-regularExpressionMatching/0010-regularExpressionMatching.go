package main

import (
	"fmt"
	"reflect"
)

func isMatch(s string, p string) bool {

	valid := true
	invalid := false

	if len(s) <= 0 || len(s) > 20 {
		return invalid
	}

	if len(p) <= 0 || len(p) > 20 {
		return invalid
	}

	// index for s
	i := 0

	// index for p
	j := 0

	for i = 0; i < len(s); i++ {
		fmt.Print("starting loop ", i, "\n")

		if p[j] == '.' {
			j += 1
			fmt.Print("continue on p[j] = .\n")
			continue
		}

		if i != 0 {

			if p[j] == '*' {

				if p[i-1] == '.' {
					j++
					fmt.Print("continue on p[j] = * followed by .\n")
					continue
				}

				if s[i-1] == s[i] {
					fmt.Print("continue s[", i-1, "] == s[", i, "]\n")
					continue
				}
			}

		}

		if s[i] == p[j] {
			fmt.Print("continue on s[", i, "] == p[", j, "]\n")

			if (j + 1) >= len(p) {
				continue
			}

			j++
			continue
		}

	}

	// if it gets here, input is valid
	return valid
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

		fmt.Printf("Test %d STARTING\n", i+1)

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
