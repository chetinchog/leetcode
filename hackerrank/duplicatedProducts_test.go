package hackerrank

import (
	"fmt"
	"testing"
)

func testDP(t *testing.T, expected int, names []string, prices []int, weights []int) {
	result := DuplicatedProducts(names, prices, weights)
	if result != expected {
		fmt.Println()
		t.Error("Expected", expected, "- Got", result)
	}
}

func TestDuplicatedProducts_1(t *testing.T) {
	fmt.Println("TestDuplicatedProducts()")
	testDP(t, 4,
		[]string{"a", "a", "a", "b", "a", "b", "a"},
		[]int{1, 1, 1, 2, 1, 2, 1},
		[]int{1, 2, 1, 1, 2, 1, 2},
	)
}
