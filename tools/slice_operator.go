package tools

//UniqueAppend append slice without the same element
func UniqueAppend(arr []interface{}, value interface{}) []interface{} {
	for _, v := range arr {
		if v == value {
			return arr
		}
	}
	arr = append(arr, value)
	return arr
}

//UniAppend append slice without the same element
func UniAppend(arr []uint, value uint) []uint {
	for _, v := range arr {
		if v == value {
			return arr
		}
	}
	arr = append(arr, value)
	return arr
}
