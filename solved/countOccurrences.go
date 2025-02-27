package solved

import (
	"fmt"
)

func countOccurrences(slice []int) map[int]int {
	result := make(map[int]int)

	for _, elem := range slice {
		result[elem]++
	}

	return result
}

func testCountOccurrences() {
	tests := [][]int{{1, 2, 3, 4}, {1, 2, 2, 3}}

	for _, t := range tests {
		fmt.Println(countOccurrences(t))
	}
}
