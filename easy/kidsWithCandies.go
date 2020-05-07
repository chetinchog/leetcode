package easy

/*
 * 1431. Kids With the Greatest Number of Candies
 * https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
 */

//KidsWithCandies - Runetime: 0 ms - Memory: 2.3 MB
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
