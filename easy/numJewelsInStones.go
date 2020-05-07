package easy

/*
 * 771. Jewels and Stones
 * https://leetcode.com/problems/jewels-and-stones/
 */

//NumJewelsInStones - Runetime: 0 ms - Memory: 2 MB
func NumJewelsInStones(J string, S string) int {
	n := 0
	for _, t := range J {
		for _, s := range S {
			if t == s {
				n++
			}
		}
	}
	return n
}
