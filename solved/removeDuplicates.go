package solved

import (
	"fmt"
)

func removeDuplicates(slice []int) []int {
	result := []int{}
	myMap := make(map[int]int)
	for _, elem := range slice {
		_, ok := myMap[elem]
		if !ok {
			result = append(result, elem)
		}
		myMap[elem]++
	}

	return result
}

func testRemoveDuplicates() {
	tests := [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 2, 5, -2, -2}}

	for _, test := range tests {
		fmt.Println(test)
		fmt.Println(removeDuplicates(test))
	}
}
