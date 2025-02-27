func moveZeroes(nums []int) {
	idx := 0

	for _, elem := range nums {
		if elem != 0 {
			nums[idx] = elem
			idx++
		}
	}

	for ; idx < len(nums); idx++ {
		nums[idx] = 0
	}

}