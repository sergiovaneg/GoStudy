package p0287

func FindDuplicate(nums []int) int {
	// main phase: search successive powers of two
	power, lambda := 1, 1
	tortoise := nums[0]
	hare := nums[tortoise] // f(x0) is the element/node next to x0.
	for tortoise != hare {
		if power == lambda {
			tortoise = hare
			power *= 2
			lambda = 0
		}
		hare = nums[hare]
		lambda += 1
	}

	// Find the position of the first repetition of length λ
	tortoise = nums[0]
	hare = nums[0]
	for i := 0; i < lambda; i++ {
		hare = nums[hare] // The distance between the hare and tortoise is now λ
	}

	// Next, the hare and tortoise move at same speed until they agree
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	return tortoise
}
