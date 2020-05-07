package main

import (
	"fmt"
	"time"

	"github.com/chetinchog/leetcode-go/easy"
)

func main() {
	fmt.Println("Working...")
	start := time.Now()
	loops := 1
	for i := 0; i < loops; i++ {
		// r := easy.KidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
		// r := easy.NumberOfSteps(123)
		// r := easy.NumJewelsInStones("aA", "aAAbbbb")
		r := easy.TwoSum([]int{2, 7, 11, 2, 7, 11, 15}, 26)
		if (i + 1) == loops {
			fmt.Println("Result:", r)
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Done in %s!", elapsed)
}
