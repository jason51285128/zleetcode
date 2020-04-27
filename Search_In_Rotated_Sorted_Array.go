package zleetcode

func searchInRotatedSortedArray(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			return mid
		}
		if a[mid] > a[start] {
			if target >= a[start] && target < a[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > a[mid] && target <= a[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
