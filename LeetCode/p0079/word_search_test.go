package p0079_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0079"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	var word string

	word = "ABCCED"
	if !p0079.Exist(board, word) {
		t.Fatalf("Expected True; got False.")
	}

	word = "SEE"
	if !p0079.Exist(board, word) {
		t.Fatalf("Expected True; got False.")
	}

	word = "ABCB"
	if p0079.Exist(board, word) {
		t.Fatalf("Expected False; got True.")
	}
}
