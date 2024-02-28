package search

func search(A []int, target int) int {
	lo := 0
	hi := len(A) - 1

	for lo <= hi {

		mid := lo + (hi-lo)/2

		if mid < 0 {
			return -1
		}

		if target < A[mid] {
			hi = mid - 1
		} else if target > A[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
