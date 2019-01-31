package tools

import "math"

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
func pager(arr []interface{}, per, split int) [][]interface{} {
	response := [][]interface{}{}
	len := len(arr)
	if split > 0 {
		per = len / split
	}
	pages := int(math.Ceil(float64(len) / float64(per)))
	for i := 0; i < pages; i++ {
		s, e := i*per, i*per+per
		if e >= len {
			e = len
		}
		parameter := arr[s:e]
		response = append(response, parameter)
	}
	return response
}
