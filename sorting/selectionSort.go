package sorting

func SelectionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := 0
		key = arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > key {
				// arr[j+1] = arr[j]
				// arr[j] = key
				arr[j+1], arr[j] = arr[j], arr[j+1]
			} else {
				break
			}
		}
	}
	return arr
}
