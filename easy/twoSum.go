package easy

//TwoSum (nums []int, target int)
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
