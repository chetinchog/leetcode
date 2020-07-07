package hackerrank

import (
	"fmt"
	"testing"
)

func areSame(a []string, b []string) bool {
	if len(a) < 1 || len(a) != len(b) {
		return false
	}
	for i, char := range a {
		if char != b[i] {
			return false
		}
	}
	return true
}

func testBV(t *testing.T, expected []string, values []string) {
	result := BracesValidation(values)
	if !areSame(result, expected) {
		fmt.Println()
		t.Error("Expected", expected, "- Got", result)
	}
}

func TestBracesValidation_1(t *testing.T) {
	fmt.Println("TestBracesValidation(\"({})\", \"{}[]()[\")")
	testBV(t, []string{"YES", "NO"}, []string{"({})", "{}[]()["})
}
