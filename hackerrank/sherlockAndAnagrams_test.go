package hackerrank

import (
	"fmt"
	"testing"
)

func test(t *testing.T, dataS string, expected int) {
	result := SherlockAndAnagrams(dataS)
	if result != int32(expected) {
		fmt.Println()
		t.Error("Expected", expected, "- Got", result)
	}
}

func TestSherlockAndAnagrams_abba(t *testing.T) {
	fmt.Println("TestSherlockAndAnagrams (abba)")
	test(t, "abba", 4)
}

func TestSherlockAndAnagrams_abcd(t *testing.T) {
	fmt.Println("TestSherlockAndAnagrams (abcd)")
	test(t, "abcd", 0)
}

func TestSherlockAndAnagrams_ifailuhkqq(t *testing.T) {
	fmt.Println("TestSherlockAndAnagrams (ifailuhkqq)")
	test(t, "ifailuhkqq", 3)
}

func TestSherlockAndAnagrams_kkkk(t *testing.T) {
	fmt.Println("TestSherlockAndAnagrams (kkkk)")
	test(t, "kkkk", 10)
}

func TestSherlockAndAnagrams_cdcd(t *testing.T) {
	fmt.Println("TestSherlockAndAnagrams (cdcd)")
	test(t, "cdcd", 10)
}
