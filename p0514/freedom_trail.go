package p0514

func leftDist(ring string, target byte, idx int) [2]int {
	dist := 0
	for ring[idx] != target {
		dist++
		if idx == 0 {
			idx = len(ring) - 1
		} else {
			idx--
		}
	}
	return [2]int{dist, idx}
}

func rightDist(ring string, target byte, idx int) [2]int {
	dist := 0
	for ring[idx] != target {
		dist++
		if idx == len(ring)-1 {
			idx = 0
		} else {
			idx++
		}
	}
	return [2]int{dist, idx}
}

func bufferDist(ring, key string,
	ring_idx, distance, pressed int) [3]int {
	if pressed == len(key) {
		return [3]int{distance, ring_idx, pressed}
	}

	// Immediate result {distance, idx}
	next_left := leftDist(ring, key[pressed], ring_idx)
	next_right := rightDist(ring, key[pressed], ring_idx)

	// Recursive result {distance, idx, pressed}
	var res_left, res_right [3]int
	if distance+next_left[0] > len(ring) {
		res_left = [3]int{distance, ring_idx, pressed}
	} else {
		res_left = bufferDist(ring, key,
			next_left[1], distance+next_left[0], pressed+1)
	}

	if distance+next_right[0] > len(ring) {
		res_right = [3]int{distance, ring_idx, pressed}
	} else {
		res_right = bufferDist(ring, key,
			next_right[1], distance+next_right[0], pressed+1)
	}

	if res_left[2] == res_right[2] { // Same amount of keys pressed
		if res_left[0] < res_right[0] {
			return res_left
		} else {
			return res_right
		}
	} else {
		if res_left[2] > res_right[2] { // More keys pressed
			return res_left
		} else {
			return res_right
		}
	}
}

func FindRotateSteps(ring string, key string) int {
	steps, ring_idx := 0, 0

	for len(key) > 0 {
		res := bufferDist(ring, key, ring_idx, 0, 0)

		steps += res[0] + res[2]
		key = key[res[2]:]
		ring_idx = res[1]
	}

	return steps
}
