package solved

import "fmt"

type TestCommonElements struct {
	testInput1, testInput2 []int
}

func convertToMap(slice []int) map[int]int {
	result := make(map[int]int)

	for _, elem := range slice {
		result[elem]++
	}
	return result
}

func commonElements(slice1 []int, slice2 []int) []int {
	map2 := convertToMap(slice2)
	result := []int{}

	for _, elem := range slice1 {
		_, ok := map2[elem]
		if ok {
			result = append(result, elem)
		}
	}

	return result
}

func testCommonElements() {
	tests := []TestCommonElements{{[]int{1, 2, 3, 4}, []int{5, 3, 4, -1}}, {[]int{4, -5, 2, 3}, []int{3, -4, -5, 12}}}

	for _, t := range tests {
		fmt.Println(commonElements(t.testInput1, t.testInput2))
	}
}
