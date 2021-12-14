package binsearch

func BinSearch(s []int, lookingFor int) int {
	if len(s) < 1 {
		return -1
	}
	low, high := 0, len(s)-1

	for low <= high {
		mid := (low + high) / 2
		midVal := s[mid]
		if midVal == lookingFor {
			return mid
		}
		if midVal > lookingFor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
