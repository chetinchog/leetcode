package easy

/*
 * 1. Two Sum
 * https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
 */

//TwoSum - Runetime: 20 ms - Memory: 2.9 MB
func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		x := target - v
		n := i + 1
		for j, y := range nums[n:] {
			if y == x {
				return []int{i, n + j}
			}
		}
	}
	return []int{}
}
