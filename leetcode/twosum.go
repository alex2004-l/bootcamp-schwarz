func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	result := []int{}

	for idx, elem := range nums {
		_, ok := myMap[target-elem]
		if ok {
			result = append(result, idx)
			result = append(result, myMap[target-elem])
			break
		}
		myMap[elem] = idx
	}

	return result
}