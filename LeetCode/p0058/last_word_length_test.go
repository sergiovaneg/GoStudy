package p0058_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0058"
)

func TestLengthOfLastWord(t *testing.T) {
	var res int

	res = p0058.LengthOfLastWord("Hello World")
	if res != 5 {
		t.Fatalf("Expected 5; got %v", res)
	}

	res = p0058.LengthOfLastWord("   fly me   to   the moon  ")
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}

	res = p0058.LengthOfLastWord("luffy is still joyboy")
	if res != 6 {
		t.Fatalf("Expected 6; got %v", res)
	}
}
