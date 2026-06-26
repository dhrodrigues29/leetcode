package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func intToRoman(num int) string {

	invalid := ""
	currentStr := ""

	if num < 1 || num > 3999 {
		return invalid
	}

	numStr := strconv.Itoa(num)

	multiplier := 1
	for i := 1; i < len(numStr); i++ {
		multiplier *= 10
	}

	for i := 0; i < len(numStr); i++ {

		digit := int(numStr[i] - '0')
		if digit == 4 {

			if multiplier == 100 {
				currentStr = currentStr + "CD"
			}

			if multiplier == 10 {
				currentStr = currentStr + "XL"
			}

			if multiplier == 1 {
				currentStr = currentStr + "IV"
			}

			multiplier /= 10
			continue
		}

		if digit == 9 {

			if multiplier == 100 {
				currentStr = currentStr + "CM"
			}

			if multiplier == 10 {
				currentStr = currentStr + "XC"
			}

			if multiplier == 1 {
				currentStr = currentStr + "IX"
			}

			multiplier /= 10
			continue
		}
		if digit == 5 {

			if multiplier == 100 {
				currentStr = currentStr + "D"
			}

			if multiplier == 10 {
				currentStr = currentStr + "L"
			}

			if multiplier == 1 {
				currentStr = currentStr + "V"
			}

			multiplier /= 10
			continue
		}
		if digit > 5 && digit < 9 {

			if multiplier == 100 {
				add := "D"
				digitNew := digit
				for digitNew > 5 {
					add = add + "C"
					digitNew--
				}
				currentStr = currentStr + add
			}

			if multiplier == 10 {
				add := "L"
				digitNew := digit
				for digitNew > 5 {
					add = add + "X"
					digitNew--
				}
				currentStr = currentStr + add
			}

			if multiplier == 1 {
				add := "V"
				digitNew := digit
				for digitNew > 5 {
					add = add + "I"
					digitNew--
				}
				currentStr = currentStr + add
			}

			multiplier /= 10
			continue
		}
		if digit < 4 {

			if multiplier == 1000 {
				add := ""
				digitNew := digit
				for digitNew > 0 {
					add = add + "M"
					digitNew--
				}
				currentStr = currentStr + add
			}

			if multiplier == 100 {
				add := ""
				digitNew := digit
				for digitNew > 0 {
					add = add + "C"
					digitNew--
				}
				currentStr = currentStr + add
			}

			if multiplier == 10 {
				add := ""
				digitNew := digit
				for digitNew > 0 {
					add = add + "X"
					digitNew--
				}
				currentStr = currentStr + add
			}

			if multiplier == 1 {
				add := ""
				digitNew := digit
				for digitNew > 0 {
					add = add + "I"
					digitNew--
				}
				currentStr = currentStr + add
			}

			multiplier /= 10
			continue
		}

	}

	return currentStr
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
