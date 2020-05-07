package easy

/*
 * 1342. Number of Steps to Reduce a Number to Zero
 * https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
 */

//NumberOfSteps - Runetime: 0 ms - Memory: 1.9 MB
func NumberOfSteps(num int) int {
	for i := 0; true; i++ {
		if num == 0 {
			return i
		} else if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
	}
	return 0
}
