package sort

type SelectionSort struct{}

func (s SelectionSort) sort(A *[]int) {
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
