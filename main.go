package main

import (
	"fmt"
	"math"
	"time"
)

// Alg for Hackerrank Array Manipulation
func Alg(n int32, queries [][]int32) int64 {
	var max int64 = 0
	arr := make([]int64, n)
	for _, op := range queries {
		for i := op[0] - 1; int32(i) < op[1]; i++ {
			arr[i] += int64(op[2])
		}
	}
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

// Alg2 for Hackerrank Array Manipulation
func Alg2(n int32, queries [][]int32) int64 {
	var max int64 = 0
	arr := make([]int64, n+2)
	for _, op := range queries {
		k := int64(op[2])
		arr[op[0]] += k
		arr[op[1]+1] -= k
	}
	var count int64
	for _, v := range arr {
		count += v
		if count > max {
			max = count
		}
	}
	return max
}

func m35(n int32) {
	for i := 0; i <= int(n); i++ {
		mo3 := i%3 == 0
		mo5 := i%5 == 0
		switch true {
		case mo3 && mo5:
			fmt.Println("FizzBuzz")
			break
		case mo3:
			fmt.Println("Fizz")
			break
		case mo5:
			fmt.Println("Buzz")
			break
		default:
			fmt.Println(i)
			break
		}
	}
}

func main() {
	fmt.Println("Working...")

	fmt.Println()
	loops := 1

	data := make([][]int32, 0)
	n := int(math.Pow(10, 7))
	m := int(2 * math.Pow(10, 5))
	for i := 1; i <= m; i++ {
		// a := int32(rand.Float64() * float64(n))
		// b := int32(float64(a) + rand.Float64()*float64(n-int(a)))
		a := int32(1)
		b := int32(n)
		k := int32(math.Pow(10, 9))
		data = append(data, []int32{a, b, k})
	}
	data2 := make([][]int32, len(data))
	copy(data2, data)
	// L := 100000
	// start := time.Now()
	// for i := 0; i < loops; i++ {
	// 	r := Alg(int32(n), data[:L])
	// 	if (i + 1) == loops {
	// 		fmt.Println("Result:", r)
	// 	}
	// }
	// t := time.Now()
	// fmt.Printf("Done in %s!\n\n", t.Sub(start))

	start2 := time.Now()
	for i := 0; i < loops; i++ {
		r := Alg2(int32(10), [][]int32{[]int32{1, 5, 3}, []int32{4, 8, 7}, []int32{6, 9, 1}})
		if (i + 1) == loops {
			fmt.Println("Result:", r)
		}
	}
	// for i := 0; i < loops; i++ {
	// 	r := Alg2(int32(n), data)
	// 	if (i + 1) == loops {
	// 		fmt.Println("Result:", r)
	// 	}
	// }
	// for i := 0; i < loops; i++ {
	// 	r := Alg2(int32(10), [][]int32{[]int32{1, 5, 3}, []int32{4, 8, 7}, []int32{6, 9, 1}})
	// 	if (i + 1) == loops {
	// 		fmt.Println("Result:", r)
	// 	}
	// }
	t2 := time.Now()
	fmt.Printf("Done in %s!\n\n", t2.Sub(start2))
}
