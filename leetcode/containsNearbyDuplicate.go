func containsNearbyDuplicate(nums []int, k int) bool {
	myMap := make(map[int]int)
	for i, elem := range nums {
		_, ok := myMap[elem]
		if ok {
			if i-myMap[elem] <= k {
				return true
			}
		}

		myMap[elem] = i
	}

	return false
}