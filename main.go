package main

import (
	"fmt"
	"time"
)

func test(s string, n int64) int64 {
	l := len(s)
	cA := countA(s)
	N := int(n)
	cR := N / l
	rL := N - cR*l
	rcA := countA(s[:rL])
	cRcA := cR * cA
	return int64(cRcA + rcA)
}

func countA(s string) int {
	r := make(map[rune]int)
	for _, v := range s {
		r[v]++
	}
	return r['a']
}

func main() {
	fmt.Println("Working...")
	start := time.Now()
	loops := 1
	for i := 0; i < loops; i++ {
		r := test("abca", 11)
		if (i + 1) == loops {
			fmt.Println("Result:", r)
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Done in %s!", elapsed)
}
