package sort

type InsertionSort struct{}

func (s *InsertionSort) sort(A *[]int) {
	n := len(*A)

	for i := 1; i < n; i++ {

		for j := i; j > 0; j-- {

			if less((*A)[j], (*A)[j-1]) {
				swap(A, j, j-1)
			} else {
				break
			}
		}
	}
}
