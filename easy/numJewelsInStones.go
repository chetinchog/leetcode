package easy

//NumJewelsInStones (J string, S string) int
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
