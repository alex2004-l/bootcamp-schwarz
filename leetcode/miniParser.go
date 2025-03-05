
import "strconv"

func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		val, _ := strconv.Atoi(s)
		result := &NestedInteger{}
		result.SetInteger(val)
		return result
	}

	stack := []*NestedInteger{}
	var current *NestedInteger
	number := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '[' {
			if current != nil {
				stack = append(stack, current)
			}
			current = &NestedInteger{}
		} else if char == ']' || char == ',' {
			if len(number) > 0 {
				val, _ := strconv.Atoi(string(number))
				num := &NestedInteger{}
				num.SetInteger(val)
				current.Add(*num)
				number = number[:0]
			}
			if char == ']' {
				if len(stack) > 0 {
					prev := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					prev.Add(*current)
					current = prev
				}
			}
		} else {
			number = append(number, char)
		}
	}

	return current
}