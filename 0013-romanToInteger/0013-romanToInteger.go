package main

import (
	"fmt"
	"reflect"
)

func romanToInt(s string) int {

}

type testCase struct {
	input    string
	expected int
}

func runTests() {

	tests := []testCase{

		// Examples
		{
			input:    "III",
			expected: 3,
		},
		{
			input:    "LVIII",
			expected: 58,
		},
		{
			input:    "MCMXCIV",
			expected: 1994,
		},

		// Small values
		{
			input:    "I",
			expected: 1,
		},
		{
			input:    "II",
			expected: 2,
		},
		{
			input:    "IV",
			expected: 4,
		},
		{
			input:    "V",
			expected: 5,
		},
		{
			input:    "VIII",
			expected: 8,
		},
		{
			input:    "IX",
			expected: 9,
		},

		// Tens
		{
			input:    "X",
			expected: 10,
		},
		{
			input:    "XIV",
			expected: 14,
		},
		{
			input:    "XL",
			expected: 40,
		},
		{
			input:    "XLIX",
			expected: 49,
		},
		{
			input:    "XCIX",
			expected: 99,
		},

		// Hundreds
		{
			input:    "CD",
			expected: 400,
		},
		{
			input:    "CMXLIV",
			expected: 944,
		},
		{
			input:    "CMXCIX",
			expected: 999,
		},

		// Thousands
		{
			input:    "MCMLXXXVII",
			expected: 1987,
		},
		{
			input:    "MMXXIV",
			expected: 2024,
		},
		{
			input:    "MMMCMXCIX",
			expected: 3999,
		},
	}

	for i, tc := range tests {

		fmt.Printf("Test %d STARTING\n", i+1)

		result := romanToInt(tc.input)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input:    %q\n", tc.input)
			fmt.Printf("Expected: %d\n", tc.expected)
			fmt.Printf("Got:      %d\n", result)
		}
	}
}

func main() {
	runTests()
}
