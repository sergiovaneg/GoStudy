package p0079

import "strings"

func findNext(board [][]byte, used [][]bool, m, n int,
	word string, word_idx int) bool {
	if board[m][n] != word[word_idx] {
		return false
	} else if word_idx == len(word)-1 {
		return true
	}

	used[m][n] = true

	// Counter-Clockwise starting from above
	if m > 0 && !used[m-1][n] { // Top
		if findNext(board, used, m-1, n, word, word_idx+1) {
			return true
		}
	}
	if n > 0 && !used[m][n-1] { // Left
		if findNext(board, used, m, n-1, word, word_idx+1) {
			return true
		}
	}
	if m < len(used)-1 && !used[m+1][n] { // Bottom
		if findNext(board, used, m+1, n, word, word_idx+1) {
			return true
		}
	}
	if n < len(used[m])-1 && !used[m][n+1] { // Right
		if findNext(board, used, m, n+1, word, word_idx+1) {
			return true
		}
	}

	used[m][n] = false
	return false
}

func Exist(board [][]byte, word string) bool {
	// Indicator matrix init
	used := make([][]bool, len(board))
	for m, m_lim := 0, len(board); m < m_lim; m++ {
		used[m] = make([]bool, len(board[m]))
		for n, n_lim := 0, len(board[m]); n < n_lim; n++ {
			used[m][n] = !strings.Contains(word, string(board[m][n]))
		}
	}

	// Matrix exploration
	for m, m_lim := 0, len(board); m < m_lim; m++ {
		for n, n_lim := 0, len(board[m]); n < n_lim; n++ {
			if used[m][n] {
				continue
			}
			if findNext(board, used, m, n, word, 0) {
				return true
			}
		}
	}

	// Default return
	return false
}
