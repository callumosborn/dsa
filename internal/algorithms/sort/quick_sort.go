package sort

type QuickSort struct{}

func (s QuickSort) sort(A *[]int) {
	var qsort func(start, end int)

	var partition = func(start, end, pivot int) int {
		if start == end {
			return start
		}

		swap(A, pivot, start)

		i := start
		j := end

		for {

			for i < end && (*A)[start] >= (*A)[i] {
				i++
			}

			for j > start && (*A)[start] <= (*A)[j] {
				j--
			}

			if i >= j {
				break
			}

			swap(A, i, j)
		}

		swap(A, start, j)

		return j
	}

	qsort = func(start, end int) {
		if end <= start {
			return
		}

		pivot := start
		spivot := partition(start, end, pivot)

		qsort(start, spivot-1)
		qsort(spivot+1, end)
	}

	qsort(0, len(*A)-1)
}
