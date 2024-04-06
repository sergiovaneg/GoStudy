package p1249_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p1249"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	var res string

	res = p1249.MinRemoveToMakeValid("lee(t(c)o)de)")
	if res != "lee(t(c)o)de" {
		t.Fatalf("Expected 'lee(t(c)o)de'; got '%v'", res)
	}

	res = p1249.MinRemoveToMakeValid("a)b(c)d")
	if res != "ab(c)d" {
		t.Fatalf("Expected 'ab(c)d'; got '%v'", res)
	}

	res = p1249.MinRemoveToMakeValid("))((")
	if res != "" {
		t.Fatalf("Expected empty string; got '%v'", res)
	}
}
