package p0930

func NumSubarraysWithSum(nums []int, goal int) int {
	c := make(chan int, 2)
	defer close(c)

	go atMostAsync(nums, goal, c)
	go atMostAsync(nums, goal-1, c)
	n1, n2 := <-c, <-c

	if n1 > n2 {
		return n1 - n2
	} else {
		return n2 - n1
	}
}

func atMostAsync(nums []int, goal int, output chan<- int) {
	var idx_0, idx_1 = 0, 0
	acc, count := 0, 0

	for idx_1 = 0; idx_1 < len(nums); idx_1++ {
		acc += nums[idx_1]
		for (acc > goal) && (idx_0 <= idx_1) {
			acc -= nums[idx_0]
			idx_0++
		}
		count += idx_1 - idx_0 + 1
	}

	output <- count
}
