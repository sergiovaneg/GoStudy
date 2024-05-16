package p0079

import (
	"strings"
	"sync"
)

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
	m_size, n_size := len(board), len(board[0])
	used := make([][][]bool, n_size)
	used[0] = make([][]bool, m_size)

	var wg sync.WaitGroup

	wg.Add(m_size)
	for m := 0; m < m_size; m++ {
		m := m
		go func() {
			defer wg.Done()
			used[0][m] = make([]bool, n_size)
			for n := 0; n < n_size; n++ {
				used[0][m][n] = !strings.Contains(word, string(board[m][n]))
			}
		}()
	}
	wg.Wait()

	wg.Add(n_size - 1)
	for n := 1; n < n_size; n++ {
		n := n
		go func() {
			defer wg.Done()
			used[n] = make([][]bool, m_size)
			wg.Add(m_size)
			for m := 0; m < m_size; m++ {
				n := n
				m := m
				go func() {
					defer wg.Done()
					used[n][m] = make([]bool, n_size)
					copy(used[n][m], used[0][m])
				}()
			}
		}()
	}

	wg.Wait()

	// Matrix exploration
	for m := 0; m < m_size; m++ {
		success := make(chan bool, n_size)
		wg.Add(n_size)
		for n := 0; n < n_size; n++ {
			n := n
			go func() {
				defer wg.Done()
				if used[n][m][n] {
					success <- false
					return
				}
				success <- findNext(board, used[n], m, n, word, 0)
			}()
		}
		wg.Wait()
		for n := 0; n < n_size; n++ {
			result := <-success
			if result {
				return true
			}
		}
	}

	// Default return
	return false
}
