package p2073

func TimeRequiredToBuy(tickets []int, k int) int {
	t := 0
	l := len(tickets)

	for tickets[k] > 0 {
		for idx := 0; idx < l; idx++ {
			if tickets[idx] == 0 {
				continue
			}
			tickets[idx]--
			t++
			if idx == k && tickets[k] == 0 {
				break
			}
		}
	}

	return t
}
