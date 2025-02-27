func plusOne(digits []int) []int {
	idx := len(digits) - 1
	digits[idx]++

	for digits[idx] > 9 && idx > 0 {
		digits[idx] -= 10
		idx--
		digits[idx] += 1
	}

	if digits[0] > 9 {
		head := []int{1}
		digits[0] -= 10
		merged := append(head, digits...)
		return merged
	}

	return digits
}