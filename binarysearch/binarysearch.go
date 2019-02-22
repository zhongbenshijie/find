package binarysearch

func BinarySearch(intSlice []int, target int) int {
	start := 0
	end := len(intSlice) - 1

	// 必须 <=
	for start <= end {
		// 此处向下取整, 也可向上取整
		mid := (start + end) / 2
		if intSlice[mid] == target {
			return mid
		} else if intSlice[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
