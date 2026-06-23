package main

import (
	"fmt"
	"reflect"
)

func convert(s string, numRows int) string {

	if numRows == 1 || len(s) == 1 {
		return s
	}

	rows := make([]string, numRows)

	currentRow := 0
	goingDown := true

	for i := 0; i < len(s); i++ {

		rows[currentRow] += string(s[i])

		if currentRow == numRows-1 {
			goingDown = false
		} else if currentRow == 0 {
			goingDown = true
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	result := ""

	for _, row := range rows {
		result += row
	}

	return result
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
