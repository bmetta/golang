package main

import "fmt"

func Merge(left []int, right []int) []int {
	data := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			data = append(data, left[0])
			left = left[1:]
		} else {
			data = append(data, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		data = append(data, left...)
	} else {
		data = append(data, right...)
	}
	return data
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	var left []int
	done := make(chan bool)
	go func() {
		left = MergeSort(data[:mid])
		done <- true
	}()
	right := MergeSort(data[mid:])
	<-done // blocks untill above go routine finish execution

	return Merge(left, right)
}

func main() {
	data := []int{8, 5, 4, 1, 9, 6, 7, 3, 2, 0}
	fmt.Println(data)
	sort_data := MergeSort(data)
	fmt.Println(sort_data)
}
