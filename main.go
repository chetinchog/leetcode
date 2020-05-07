package main

import (
	"fmt"
	"time"
)

func twoSum(nums []int, target int) []int {
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

func kidsWithCandies(candies []int, extraCandies int) []bool {
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

func numberOfSteps(num int) int {
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

func numJewelsInStones(J string, S string) int {
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

func main() {
	fmt.Println("Working...")
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		// for i := 0; i < 1; i++ {
		// twoSum([]int{2, 7, 11, 2, 7, 11, 15}, 26)
		// kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
		// numberOfSteps(123)
		numJewelsInStones("aA", "aAAbbbb")
	}
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Done in %s!", elapsed)
}
