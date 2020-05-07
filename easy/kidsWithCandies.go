package easy

//KidsWithCandies (candies []int, extraCandies int) []bool
func KidsWithCandies(candies []int, extraCandies int) []bool {
	r := []bool{}
	for _, v := range candies {
		s := true
		e := v + extraCandies
		for _, y := range candies {
			if e < y {
				s = false
			}
		}
		r = append(r, s)
	}
	return r
}
