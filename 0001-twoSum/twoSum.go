package main

import (
    "fmt"
    "reflect"
)

// Mmy solution
func twoSum(nums []int, target int) []int {
    // make map to save seen values
    seen := make(map[int]int)

    // for all indices in nums
    for i, x := range nums {
        // needed is the value needed to get to target when added to the current x
        needed := target - x

        // check if need value was already seen
        if idx, found := seen[needed]; found {
            fmt.Println([]int{idx, i})
            return []int{idx, i}
        }

        // if not store current number
        seen[x] = i
    }
    // if no solution exists
    fmt.Println("There is no possible sum equal to", target, "in the array", nums)
    return nil
}


// 2. Test structure
type testCase struct {
    nums     []int
    target   int
    expected []int
}

// 3. Runner
func runTests() {
    tests := []testCase{
        {nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
        {nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
        {nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
        {nums: []int{2, 7, 11, 15}, target: 27, expected: nil},
    }

    for i, tc := range tests {
        result := twoSum(tc.nums, tc.target)

        if reflect.DeepEqual(result, tc.expected) {
            fmt.Printf("Test %d PASSED \n", i+1)
        } else {
            fmt.Printf("Test %d FAILED\n", i+1)
            fmt.Printf("  Input: nums=%v target=%d\n", tc.nums, tc.target)
            fmt.Printf("  Expected: %v\n", tc.expected)
            fmt.Printf("  Got: %v\n", result)
        }
    }
}

// 4. Main
func main() {
    runTests()
}