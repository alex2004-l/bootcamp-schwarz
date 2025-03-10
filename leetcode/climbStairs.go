func climbStairs(n int) int {
	result := []int{1, 2}

	for i := 2; i < n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}

	return result[n-1]
}