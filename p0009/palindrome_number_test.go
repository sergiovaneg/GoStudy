package p0009_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0009"
)

func TestIsPalindrome(t *testing.T) {
	if !p0009.IsPalindrome(121) {
		t.Fatal("Expected true; got false")
	}

	if p0009.IsPalindrome(-121) {
		t.Fatal("Expected false; got true")
	}

	if p0009.IsPalindrome(10) {
		t.Fatal("Expected false; got true")
	}
}
