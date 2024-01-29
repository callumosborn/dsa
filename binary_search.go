package main

// search performs binary search on a sorted integer slice.
//
// It returns the index of the target element if found, otherwise returns -1.
//
// The function includes a check for integer overflow during the mid calculation
// to ensure the correctness and safety of the binary search algorithm.
//
// Parameters:
//   - A: A sorted integer slice to search.
//   - target: The target integer to search for in the slice.
//
// Returns:
//   - int: The index of the target element if found, otherwise -1.
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
