package main

func merge(arr1, arr2 []int) []int {
	var sortedArr []int
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			sortedArr = append(sortedArr, arr1[0])
			arr1 = arr1[1:]
		} else {
			sortedArr = append(sortedArr, arr2[0])
			arr2 = arr2[1:]
		}
	}
	sortedArr = append(sortedArr, arr1...)
	sortedArr = append(sortedArr, arr2...)
	return sortedArr
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var mid int = len(arr) / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
