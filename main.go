package main

import "fmt"

// TwoSum returns indices i < j with nums[i]+nums[j] == target, or nil when
// no pair exists. One pass with a value→index map: each element looks up
// the complement it needs among elements already seen.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for j, v := range nums {
		if i, ok := seen[target-v]; ok {
			return []int{i, j}
		}
		seen[v] = j
	}
	return nil
}

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
}
