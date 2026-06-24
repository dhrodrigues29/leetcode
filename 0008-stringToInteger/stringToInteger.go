package main

import (
	"fmt"
	"reflect"
)

func myAtoi(s string) int {

	i := 0
	result := 0
	isNegative := false
	charSeen := false

	for i < len(s) {

		if charSeen == false {

			//spaces only matters with no char have been seen yet
			if s[i] == ' ' {
				i++
				continue
			}

			if s[i] == '+' {
				charSeen = true
				i++
				continue
			}

			if s[i] == '-' {
				charSeen = true
				isNegative = true
				i++
				continue
			}

		}

		if s[i] < '0' || s[i] > '9' {
			if isNegative {
				return -result
			}
			return result
		}

		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')

			if !isNegative {
				if result > 214748364 ||
					(result == 214748364 && digit > 7) {
					return 2147483647
				}
			} else {
				if result > 214748364 ||
					(result == 214748364 && digit > 8) {
					return -2147483648
				}
			}

			result = result*10 + digit
			charSeen = true
		}

		i++

	}

	if isNegative {
		result = -result
	}

	return result
}

type testCase struct {
	input    string
	expected int
}

func runTests() {

	tests := []testCase{
		{
			input:    "42",
			expected: 42,
		},
		{
			input:    "   -042",
			expected: -42,
		},
		{
			input:    "1337c0d3",
			expected: 1337,
		},
		{
			input:    "0-1",
			expected: 0,
		},
		{
			input:    "words and 987",
			expected: 0,
		},
		{
			input:    "+1",
			expected: 1,
		},
		{
			input:    "-1",
			expected: -1,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "     ",
			expected: 0,
		},
		{
			input:    "00000123",
			expected: 123,
		},
		{
			input:    "   +123",
			expected: 123,
		},
		{
			input:    "4193 with words",
			expected: 4193,
		},
		{
			input:    "words 123",
			expected: 0,
		},
		{
			input:    "+-12",
			expected: 0,
		},
		{
			input:    "-+12",
			expected: 0,
		},
		{
			input:    "3.14159",
			expected: 3,
		},
		{
			input:    "2147483647",
			expected: 2147483647,
		},
		{
			input:    "2147483648",
			expected: 2147483647,
		},
		{
			input:    "999999999999999999999999999",
			expected: 2147483647,
		},
		{
			input:    "-2147483648",
			expected: -2147483648,
		},
		{
			input:    "-2147483649",
			expected: -2147483648,
		},
		{
			input:    "-999999999999999999999999999",
			expected: -2147483648,
		},
		{
			input:    "00000abc123",
			expected: 0,
		},
		{
			input:    "-00000123",
			expected: -123,
		},
		{
			input:    "  -0012a42",
			expected: -12,
		},
	}

	for i, tc := range tests {

		result := myAtoi(tc.input)

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
