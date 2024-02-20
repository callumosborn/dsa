package algorithms

// sort is a function that sorts the elements of the given slice A in ascending order.
// The function modifies the original slice.
//
// Parameters:
// - A: A pointer to the slice of integers to be sorted.
func sort(A *[]int) {
	n := len(*A)

	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {

			if less((*A)[j], (*A)[min]) {
				min = j
			}
		}

		swap(A, i, min)
	}
}

// less compares two integers and returns true if the first integer is less than the second integer.
func less(a int, b int) bool {
	return a < b
}

// swap swaps the elements at indices a and b in the given slice A.
func swap(A *[]int, a int, b int) {
	(*A)[a], (*A)[b] = (*A)[b], (*A)[a]
}
