package hackerrank

import (
	"fmt"
	"testing"
)

func testSaA(t *testing.T, dataS string, expected int) {
	result := SherlockAndAnagrams(dataS)
	if result != int32(expected) {
		fmt.Println()
		t.Error("Expected", expected, "- Got", result)
	}
}

// func TestSherlockAndAnagrams_abba(t *testing.T) {
// 	fmt.Println("TestSherlockAndAnagrams (abba)")
// 	testSaA(t, "abba", 4)
// }

// func TestSherlockAndAnagrams_abcd(t *testing.T) {
// 	fmt.Println("TestSherlockAndAnagrams (abcd)")
// 	testSaA(t, "abcd", 0)
// }

// func TestSherlockAndAnagrams_ifailuhkqq(t *testing.T) {
// 	fmt.Println("TestSherlockAndAnagrams (ifailuhkqq)")
// 	testSaA(t, "ifailuhkqq", 3)
// }

// func TestSherlockAndAnagrams_kkkk(t *testing.T) {
// 	fmt.Println("TestSherlockAndAnagrams (kkkk)")
// 	testSaA(t, "kkkk", 10)
// }

// func TestSherlockAndAnagrams_cdcd(t *testing.T) {
// 	fmt.Println("TestSherlockAndAnagrams (cdcd)")
// 	testSaA(t, "cdcd", 10)
// }
