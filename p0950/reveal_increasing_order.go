package p0950

import (
	"slices"
)

/* Original version with memory reservation only
func roll_and_prepend(sorted_deck, new_deck []int) []int {
	l_sorted := len(sorted_deck)
	l_new := len(new_deck)
	if l_new != 0 {
		new_deck = append(
			[]int{sorted_deck[l_sorted-1], new_deck[l_new-1]},
			new_deck[:l_new-1]...)
	} else {
		new_deck = []int{sorted_deck[l_sorted-1]}
	}

	if l_sorted == 1 {
		return new_deck
	} else {
		return roll_and_prepend(sorted_deck[:l_sorted-1], new_deck)
	}
}
*/

func roll_and_prepend(
	sorted_deck, new_deck []int,
	idx int) []int {
	aux := new_deck[idx-1]
	copy(new_deck[2:], new_deck[:idx-1])
	new_deck[1] = aux

	new_deck[0] = sorted_deck[0]
	if len(sorted_deck) == 1 {
		return new_deck
	} else {
		return roll_and_prepend(sorted_deck[1:], new_deck, idx+1)
	}
}

func DeckRevealedIncreasing(deck []int) []int {
	if len(deck) < 2 {
		return deck
	}
	slices.SortStableFunc(deck, func(a, b int) int { return b - a })

	new_deck := make([]int, len(deck))
	new_deck[0] = deck[0]
	return roll_and_prepend(deck[1:], new_deck, 1)
}
