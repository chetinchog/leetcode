package main

import (
	"fmt"
	"math"
	"testing"
)

// func TestAlg(t *testing.T) {
// 	fmt.Println("Test Alg")
// 	data := make([][]int32, 0)
// 	n := int32(math.Pow(10, 7))
// 	m := int(2 * math.Pow(10, 5))
// 	for i := 1; i <= m; i++ {
// 		a := int32(1)
// 		b := int32(n)
// 		k := int32(math.Pow(10, 9))
// 		data = append(data, []int32{a, b, k})
// 	}
// 	var expected int64 = 400000000000
// 	L := 400
// 	result := Alg(n, data[:L])
// 	if expected != result {
// 		t.Error("Failed")
// 	}
// }

// func TestAlg2(t *testing.T) {
// 	fmt.Println("Test Alg2")
// 	data := make([][]int32, 0)
// 	n := int32(math.Pow(10, 7))
// 	m := int(2 * math.Pow(10, 5))
// 	for i := 1; i <= m; i++ {
// 		a := int32(1)
// 		b := int32(n)
// 		k := int32(math.Pow(10, 9))
// 		data = append(data, []int32{a, b, k})
// 	}
// 	var expected int64 = 400000000000
// 	L := 400
// 	result := Alg2(n, data[:L])
// 	if expected != result {
// 		fmt.Println(expected, result)
// 		t.Error("Failed")
// 	}
// }

var expected int64 = 12

func TestAlg(t *testing.T) {
	fmt.Println("TestAlg")
	result := Alg(int32(10), [][]int32{[]int32{1, 5, 3}, []int32{4, 8, 7}, []int32{6, 9, 1}, []int32{10, 10, 12}})
	if expected != result {
		fmt.Println("expected", expected, "result", result)
		t.Error("Failed")
	}
}

func TestAlg2(t *testing.T) {
	fmt.Println("TestAlg2")
	result := Alg2(int32(10), [][]int32{[]int32{1, 5, 3}, []int32{4, 8, 7}, []int32{6, 9, 1}, []int32{10, 10, 12}})
	if expected != result {
		fmt.Println("expected", expected, "result", result)
		t.Error("Failed")
	}
}

var L int = 1000

func BenchmarkAlg(b *testing.B) {
	fmt.Println("Benchmark Alg")
	data := make([][]int32, 0)
	n := int32(math.Pow(10, 7))
	m := int(2 * math.Pow(10, 5))
	for i := 1; i <= m; i++ {
		a := int32(1)
		b := int32(n)
		k := int32(math.Pow(10, 9))
		data = append(data, []int32{a, b, k})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Alg(n, data[:L])
	}
}

func BenchmarkAlg2(b *testing.B) {
	fmt.Println("Benchmark Alg2")
	data := make([][]int32, 0)
	n := int32(math.Pow(10, 7))
	m := int(2 * math.Pow(10, 5))
	for i := 1; i <= m; i++ {
		a := int32(1)
		b := int32(n)
		k := int32(math.Pow(10, 9))
		data = append(data, []int32{a, b, k})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Alg2(n, data[:L])
	}
}
