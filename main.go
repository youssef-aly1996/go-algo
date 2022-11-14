package main

import (
	"algo-lab/examples"
	"algo-lab/searching"
	"algo-lab/sorting"
	"fmt"
)

func main() {
	arr := []int{39, 28, -44, -4, 10, 83, -11}
	sorted := sorting.MergeSort(arr)
	fmt.Println(sorted)
	index := searching.BinarySearch(sorted, 4)
	fmt.Println(index)
	segrated := examples.MergeSeg(arr)
	fmt.Println(segrated)

}


