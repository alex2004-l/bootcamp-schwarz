package solved

import "fmt"

type testFind struct {
	testInput []int
	output    int
}

func findIndex(slice []int, val int) []int {
	result := []int{}

	for idx, elem := range slice {
		if elem == val {
			result = append(result, idx)
		}
	}

	return result
}

func testFindIdx() {
	tests := []testFind{{[]int{1, 2, 3, 4}, 1}, {[]int{1, 2, 2, 3}, 2}}

	for _, t := range tests {
		fmt.Println(findIndex(t.testInput, t.output))
	}
}
