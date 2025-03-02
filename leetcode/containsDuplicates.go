func containsDuplicate(nums []int) bool {
    myMap := make(map[int]int)

    for _, elem := range nums {
        myMap[elem] ++
    }

    for _ , val := range myMap {
        if val > 1 {
            return true
        }
    }

    return false
}