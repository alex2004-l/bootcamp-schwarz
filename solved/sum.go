package solved

import (
	"fmt"
)

func sum(slice []int) int {
	sum := 0
	for _, elem := range slice {
		sum += elem
	}

	return sum
}

func testSum() {
	test := []int{1, 2, 3, 4, 5}

	if sum(test) != 15 {
		fmt.Println("Incorrect answer")
	} else {
		fmt.Println("Correct answer")
	}
}
