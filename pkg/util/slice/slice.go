// Package slice slice convert tool
package slice

import "math"

// UintInSlice if uint in slice
func UintInSlice(a uint, list []uint) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// DoByPage 依照陣列批次處理任務
func DoByPage(per int, data []interface{}, fn func(onepage []interface{})) {
	times := int(math.Ceil(float64(len(data)) / float64(per)))
	for i := 0; i < times; i++ {
		sidx := i * per
		eidx := (i + 1) * per
		if eidx > len(data) {
			eidx = len(data)
		}
		tasks := data[sidx:eidx]
		fn(tasks)
	}
}

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
