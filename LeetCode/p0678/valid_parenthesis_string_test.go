package p0678_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0678"
)

func TestCheckValidString(t *testing.T) {
	if !p0678.CheckValidString("()") {
		t.Fatalf("Expected True; got False.")
	}

	if !p0678.CheckValidString("(*)") {
		t.Fatalf("Expected True; got False.")
	}

	if !p0678.CheckValidString("(*))") {
		t.Fatalf("Expected True; got False.")
	}

	if p0678.CheckValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())") {
		t.Fatalf("Expected False; got True")
	}
}
