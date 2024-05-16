package p1656_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p1656"
)

func TestOrderedStream(t *testing.T) {
	obj := p1656.Constructor(5)

	var res []string

	res = obj.Insert(3, "ccccc")
	if !reflect.DeepEqual(res, []string{}) {
		t.Fatalf("Expected empty buffer; got %v", res)
	}

	res = obj.Insert(1, "aaaaa")
	if !reflect.DeepEqual(res, []string{"aaaaa"}) {
		t.Fatalf("Expected [aaaaa]; got %v", res)
	}

	res = obj.Insert(2, "bbbbb")
	if !reflect.DeepEqual(res, []string{"bbbbb", "ccccc"}) {
		t.Fatalf("Expected [bbbbb,ccccc]; got %v", res)
	}

	res = obj.Insert(5, "eeeee")
	if !reflect.DeepEqual(res, []string{}) {
		t.Fatalf("Expected empty buffer; got %v", res)
	}

	res = obj.Insert(4, "ddddd")
	if !reflect.DeepEqual(res, []string{"ddddd", "eeeee"}) {
		t.Fatalf("Expected [ddddd,eeeee]; got %v", res)
	}
}
