package searching

func BinarySearch(arr []int, item int) int {
	start := 0
	end := len(arr) -1
	for i:= 0; i < len(arr); i++ {
		mid := (start + end) /2
		if item == arr[mid] {
			return mid
		} else if item > arr[mid] {
			start = mid + 1
		} else if item < arr[mid] {
			end = mid - 1
		}
	}
	return -1
}