package main

import (
	"fmt"
	"reflect"
	"sort"
)

func letterCombinations(digits string) []string {

}

type testCase struct {
	input    string
	expected []string
}

func normalize(result []string) []string {
	sort.Strings(result)
	return result
}

func runTests() {

	tests := []testCase{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "2",
			expected: []string{"a", "b", "c"},
		},
		{
			input:    "3",
			expected: []string{"d", "e", "f"},
		},
		{
			input:    "4",
			expected: []string{"g", "h", "i"},
		},
		{
			input:    "5",
			expected: []string{"j", "k", "l"},
		},
		{
			input:    "6",
			expected: []string{"m", "n", "o"},
		},
		{
			input:    "7",
			expected: []string{"p", "q", "r", "s"},
		},
		{
			input:    "8",
			expected: []string{"t", "u", "v"},
		},
		{
			input:    "9",
			expected: []string{"w", "x", "y", "z"},
		},
		{
			input: "23",
			expected: []string{
				"ad", "ae", "af",
				"bd", "be", "bf",
				"cd", "ce", "cf",
			},
		},
		{
			input: "79",
			expected: []string{
				"pw", "px", "py", "pz",
				"qw", "qx", "qy", "qz",
				"rw", "rx", "ry", "rz",
				"sw", "sx", "sy", "sz",
			},
		},
		{
			input: "22",
			expected: []string{
				"aa", "ab", "ac",
				"ba", "bb", "bc",
				"ca", "cb", "cc",
			},
		},
		{
			input: "27",
			expected: []string{
				"ap", "aq", "ar", "as",
				"bp", "bq", "br", "bs",
				"cp", "cq", "cr", "cs",
			},
		},
		{
			input: "92",
			expected: []string{
				"wa", "wb", "wc",
				"xa", "xb", "xc",
				"ya", "yb", "yc",
				"za", "zb", "zc",
			},
		},
		{
			input: "99",
			expected: []string{
				"ww", "wx", "wy", "wz",
				"xw", "xx", "xy", "xz",
				"yw", "yx", "yy", "yz",
				"zw", "zx", "zy", "zz",
			},
		},
		{
			input: "234",
			expected: []string{
				"adg", "adh", "adi",
				"aeg", "aeh", "aei",
				"afg", "afh", "afi",
				"bdg", "bdh", "bdi",
				"beg", "beh", "bei",
				"bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi",
				"ceg", "ceh", "cei",
				"cfg", "cfh", "cfi",
			},
		},
		{
			input: "68",
			expected: []string{
				"mt", "mu", "mv",
				"nt", "nu", "nv",
				"ot", "ou", "ov",
			},
		},
		{
			input: "78",
			expected: []string{
				"pt", "pu", "pv",
				"qt", "qu", "qv",
				"rt", "ru", "rv",
				"st", "su", "sv",
			},
		},
		{
			input: "89",
			expected: []string{
				"tw", "tx", "ty", "tz",
				"uw", "ux", "uy", "uz",
				"vw", "vx", "vy", "vz",
			},
		},
		{
			input: "2345",
			expected: []string{
				"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil",
				"aegj", "aegk", "aegl", "aehj", "aehk", "aehl", "aeij", "aeik", "aeil",
				"afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik", "afil",
				"bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil",
				"begj", "begk", "begl", "behj", "behk", "behl", "beij", "beik", "beil",
				"bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij", "bfik", "bfil",
				"cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil",
				"cegj", "cegk", "cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil",
				"cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl", "cfij", "cfik", "cfil",
			},
		},
	}

	var failed []int

	for i, tc := range tests {

		result := normalize(letterCombinations(tc.input))
		expected := normalize(tc.expected)

		if reflect.DeepEqual(result, expected) {
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
			result := normalize(letterCombinations(tc.input))
			expected := normalize(tc.expected)

			fmt.Printf("Test %d\n", i+1)
			fmt.Printf("Input:    %q\n", tc.input)
			fmt.Printf("Expected: %v\n", expected)
			fmt.Printf("Got:      %v\n\n", result)
		}
	}

	fmt.Printf("Summary: %d/%d tests passed\n",
		len(tests)-len(failed), len(tests))
}

func main() {
	runTests()
}
