package p0205_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0205"
)

func TestIsIsomorphic(t *testing.T) {
	if !p0205.IsIsomorphic("egg", "add") {
		t.Fatalf("Expected true; got false.")
	}

	if p0205.IsIsomorphic("foo", "bar") {
		t.Fatalf("Expected false; got true.")
	}

	if !p0205.IsIsomorphic("paper", "title") {
		t.Fatalf("Expected true; got false.")
	}
}
