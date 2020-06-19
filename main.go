package main

import (
	"fmt"
	"time"
)

func test(q []int32) {
	B := 0
	L := len(q)
	Q := make(map[int32]int)
	for i, v := range q {
		vi1 := int(v) - (i + 1)
		if vi1 > 2 {
			fmt.Println("Too chaotic")
			return
		}
		Q[v] = i
	}
	for i := 0; i < L-1; i++ {
		correctValue := i + 1
		for k := Q[int32(correctValue)]; k > i; k-- {
			B++
			Q[q[k-1]], Q[q[k]] = Q[q[k]], Q[q[k-1]]
			q[k-1], q[k] = q[k], q[k-1]
		}
	}
	fmt.Println("Bribes in Argentina:", B)
}

func test2(q []int32) {
	B := 0
	L := len(q)
	Q := make(map[int32]int)
	for i, v := range q {
		vi1 := int(v) - (i + 1)
		if vi1 > 2 {
			fmt.Println("Too chaotic")
			return
		}
		Q[v] = i
	}
	for i := 0; i < L-1; i++ {
		iNext := i + 1
		if int(q[iNext]) == (iNext) {
			B++
			q[i], q[iNext] = q[iNext], q[i]
		} else {
			for k := Q[int32(iNext)]; k > i; k-- {
				B++
				Q[q[k-1]], Q[q[k]] = Q[q[k]], Q[q[k-1]]
				q[k-1], q[k] = q[k], q[k-1]
			}
		}
	}
	fmt.Println("Bribes in Argentina:", B)
}

func main() {
	fmt.Println("Working...")
	fmt.Println()
	loops := 1

	var argentina []int32
	for i := 3; i <= 10000000; i++ {
		argentina = append(argentina, int32(i))
	}
	argentina = append(argentina, int32(1), int32(2))
	var argentina2 []int32
	for i := 3; i <= 10000000; i++ {
		argentina2 = append(argentina2, int32(i))
	}
	argentina2 = append(argentina2, int32(1), int32(2))

	start2 := time.Now()
	for i := 0; i < loops; i++ {
		test2(argentina2)
		if (i + 1) == loops {
			// fmt.Println("Result:", r)
		}
	}
	t2 := time.Now()
	elapsed2 := t2.Sub(start2)
	fmt.Printf("Old Alg: Done in %s!", elapsed2)

	fmt.Println()
	fmt.Println()

	start := time.Now()
	for i := 0; i < loops; i++ {
		test(argentina)
		if (i + 1) == loops {
			// fmt.Println("Result:", r)
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Old Alg: Done in %s!", elapsed)

	fmt.Println()
	fmt.Println()
	fmt.Println("Old Alg vs New Alg =", int(elapsed2)/int(elapsed))
}
