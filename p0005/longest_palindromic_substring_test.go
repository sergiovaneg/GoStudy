package p0005_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0005"
)

func TestLongestPalindrome(t *testing.T) {
	var res string

	res = p0005.LongestPalindrome("babad")
	if res != "bab" && res != "aba" {
		t.Fatalf("Expected 'bab' or 'aba'; got %v", res)
	}

	res = p0005.LongestPalindrome("cbbd")
	if res != "bb" {
		t.Fatalf("Expected 'bb'; got %v", res)
	}
}
