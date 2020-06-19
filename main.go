package main

import (
	"fmt"
	"time"
)

func algorithm(arr []int32) int32 {
	B := 0
	L := len(arr)
	Q := make(map[int32]int)
	for i, v := range arr {
		Q[v] = i
	}
	for i := 0; i < L-1; i++ {
		correctValue := i + 1
		position := Q[int32(correctValue)]
		if i != position {
			B++
			Q[arr[i]], Q[arr[position]] = Q[arr[position]], Q[arr[i]]
			arr[i], arr[position] = arr[position], arr[i]
		}
	}
	return int32(B)
}

func main() {
	fmt.Println("Working...")
	fmt.Println()
	loops := 1

	var data []int32
	for i := 1000000; i > 0; i-- {
		data = append(data, int32(i))
	}

	start := time.Now()
	for i := 0; i < loops; i++ {
		r := algorithm([]int32{4, 3, 1, 2})
		if (i + 1) == loops {
			fmt.Println("Result:", r)
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Done in %s!\n\n", elapsed)
}
