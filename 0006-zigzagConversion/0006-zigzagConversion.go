package main

import (
	"fmt"
	"reflect"
)

func convert(s string, numRows int) string {

	// TODO: implement

	return ""
}

type testCase struct {
	inputS       string
	inputNumRows int
	expected     string
}

func runTests() {

	tests := []testCase{
		{
			inputS:       "PAYPALISHIRING",
			inputNumRows: 3,
			expected:     "PAHNAPLSIIGYIR",
		},
		{
			inputS:       "PAYPALISHIRING",
			inputNumRows: 4,
			expected:     "PINALSIGYAHRPI",
		},
		{
			inputS:       "A",
			inputNumRows: 1,
			expected:     "A",
		},
		{
			inputS:       "ABCD",
			inputNumRows: 2,
			expected:     "ACBD",
		},
		{
			inputS:       "ABC",
			inputNumRows: 5,
			expected:     "ABC",
		},
		{
			inputS:       "ABCD",
			inputNumRows: 4,
			expected:     "ABCD",
		},
	}

	for i, tc := range tests {

		result := convert(tc.inputS, tc.inputNumRows)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Test %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test %d FAILED\n", i+1)
			fmt.Printf("Input S:       %v\n", tc.inputS)
			fmt.Printf("Input NumRows: %v\n", tc.inputNumRows)
			fmt.Printf("Expected:      %v\n", tc.expected)
			fmt.Printf("Got:           %v\n", result)
		}
	}
}

func main() {
	runTests()
}
