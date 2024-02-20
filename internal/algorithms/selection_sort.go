package algorithms

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

func less(a int, b int) bool {
	return a < b
}

func swap(A *[]int, a int, b int) {
	(*A)[a], (*A)[b] = (*A)[b], (*A)[a]
}
