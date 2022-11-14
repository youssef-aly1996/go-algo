package examples

func segregate(left, right []int) []int {
	var segregated []int
	var pos []int
	var neg []int
	for len(left) > 0 {
		if left[0] > 0 {
			pos = append(pos, left[0])
			left = left[1:]
		} else {
			neg = append(neg, left[0])
			left = left[1:]
		}
	}
	for len(right) > 0 {

		if right[0] > 0 {
			pos = append(pos, right[0])
			right = right[1:]
		} else {
			neg = append(neg, right[0])
			right = right[1:]
		}
	}
	segregated = append(segregated, neg...)
	segregated = append(segregated, pos...)

	return segregated
}

func MergeSeg(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var mid int = len(arr) / 2
	left := MergeSeg(arr[:mid])
	right := MergeSeg(arr[mid:])
	return segregate(left, right)
}
