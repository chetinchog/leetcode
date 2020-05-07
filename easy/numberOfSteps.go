package easy

//NumberOfSteps (num int) int
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
