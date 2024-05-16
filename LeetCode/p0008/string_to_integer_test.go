package p0008_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0008"
)

func TestMyAtoi(t *testing.T) {
	var res int

	res = p0008.MyAtoi("1337c0d3")
	if res != 1337 {
		t.Fatalf("Expected 1337; got %v", res)
	}

	res = p0008.MyAtoi("0-1")
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}

	res = p0008.MyAtoi("words and 987")
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}
}
