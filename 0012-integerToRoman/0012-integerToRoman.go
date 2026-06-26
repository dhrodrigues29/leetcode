package main

import (
	"fmt"
	"reflect"
)

func intToRoman(num int) string {

	invalid := ""

	// Constraints
	if num < 1 || num > 3999 {
		return invalid
	}

	// ======================================================
	// Your solution starts here
	//
	// Suggested ideas:
	//
	// - Keep track of the remaining value.
	// - Build the Roman numeral from left to right.
	// - Handle thousands, hundreds, tens and ones.
	// - Remember the six subtractive cases:
	//     IV  IX
	//     XL  XC
	//     CD  CM
	//
	// ======================================================

	return ""
}

type testCase struct {
	num      int
	expected string
}

func runTests() {

	tests := []testCase{

		// Examples
		{
			num:      3749,
			expected: "MMMDCCXLIX",
		},
		{
			num:      58,
			expected: "LVIII",
		},
		{
			num:      1994,
			expected: "MCMXCIV",
		},

		// Small numbers
		{
			num:      1,
			expected: "I",
		},
		{
			num:      2,
			expected: "II",
		},
		{
			num:      3,
			expected: "III",
		},
		{
			num:      4,
			expected: "IV",
		},
		{
			num:      5,
			expected: "V",
		},
		{
			num:      6,
			expected: "VI",
		},
		{
			num:      8,
			expected: "VIII",
		},
		{
			num:      9,
			expected: "IX",
		},

		// Tens
		{
			num:      10,
			expected: "X",
		},
		{
			num:      14,
			expected: "XIV",
		},
		{
			num:      19,
			expected: "XIX",
		},
		{
			num:      20,
			expected: "XX",
		},
		{
			num:      40,
			expected: "XL",
		},
		{
			num:      49,
			expected: "XLIX",
		},
		{
			num:      50,
			expected: "L",
		},
		{
			num:      90,
			expected: "XC",
		},
		{
			num:      99,
			expected: "XCIX",
		},

		// Hundreds
		{
			num:      100,
			expected: "C",
		},
		{
			num:      400,
			expected: "CD",
		},
		{
			num:      444,
			expected: "CDXLIV",
		},
		{
			num:      500,
			expected: "D",
		},
		{
			num:      900,
			expected: "CM",
		},
		{
			num:      999,
			expected: "CMXCIX",
		},

		// Thousands
		{
			num:      1000,
			expected: "M",
		},
		{
			num:      1987,
			expected: "MCMLXXXVII",
		},
		{
			num:      2024,
			expected: "MMXXIV",
		},
		{
			num:      2421,
			expected: "MMCDXXI",
		},
		{
			num:      3888,
			expected: "MMMDCCCLXXXVIII",
		},
		{
			num:      3999,
			expected: "MMMCMXCIX",
		},

		// Invalid input
		{
			num:      0,
			expected: "",
		},
		{
			num:      -10,
			expected: "",
		},
		{
			num:      4000,
			expected: "",
		},
	}

	for i, tc := range tests {

		fmt.Printf("Test %d STARTING\n", i+1)

		result := intToRoman(tc.num)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)

		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input:    %d\n", tc.num)
			fmt.Printf("Expected: %q\n", tc.expected)
			fmt.Printf("Got:      %q\n", result)
		}
	}
}

func main() {
	runTests()
}
