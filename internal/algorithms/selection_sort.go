package algorithms

func sort(A *[]int) {
	n := len(*A)

	// The reason that for first for loop we only want to iterate through n-1
	// elements in the slice is: the element in the first position is sorted
	// when we get to the end of the algorithm.
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

// less is a function that compares two integers and returns true
// if the first integer 'a' is less than the second integer 'b'.
//
// Parameters:
//   - a: An integer representing the first number to compare.
//   - b: An integer representing the second number to compare.
//
// Returns:
//   - bool: True if 'a' is less than 'b', otherwise false.
func less(a int, b int) bool {
	return a < b
}

// swap is a function that swaps the elements at two given indices in a slice of integers.
//
// Parameters:
//   - A: A pointer to a slice of integers where the elements are to be swapped.
//   - a: An integer representing the index of the first element to be swapped.
//   - b: An integer representing the index of the second element to be swapped.
func swap(A *[]int, a int, b int) {
	(*A)[a], (*A)[b] = (*A)[b], (*A)[a]
}
