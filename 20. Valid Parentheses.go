package main
func isValid(s string) bool {
	array := []byte(s)
	var left = make([]byte, 0)
	if array[0] == ']' || array[0] == '}' || array[0] == ')' {
		return false
	}
	if len(array) == 0 {
		return false
	}
	for i := 0; i < len(array); i++ {
		if array[i] == '[' || array[i] == '{' || array[i] == '(' {
			left = append(left, array[i])
		} else {
			if array[i] == ')' && left[len(left)-1] == '(' {
				left = left[:len(left)-1]
			} else if array[i] == ']' && left[len(left)-1] == '[' {
				left = left[:len(left)-1]
			} else if array[i] == '}' && left[len(left)-1] == '{' {
				left = left[:len(left)-1]
			}
		}
	}
	if len(left) == 0 {
		return true
	}
	return false
}
