package solved

import (
	"fmt"
)

func average(slice []int) float64 {
	sum := 0
	for _, elem := range slice {
		sum += elem
	}

	return float64(sum) / float64(len(slice))
}

func testAverage() {
	test := []int{1, 2, 3, 4, 5}

	if average(test) != 3.0 {
		fmt.Println("Incorrect answer")
	} else {
		fmt.Println("Correct answer")
	}
}
