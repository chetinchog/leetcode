package hackerrank

// DuplicatedProducts returns number of duplicate products.
func DuplicatedProducts(names []string, prices []int, weights []int) int {
	list := make(map[string]bool, 0)
	l := len(names)
	count := 0
	for i := 0; i < l; i++ {
		hash := names[i] + string(prices[i]) + string(weights[i])
		if list[hash] {
			count++
		} else {
			list[hash] = true
		}
	}
	return count
}
