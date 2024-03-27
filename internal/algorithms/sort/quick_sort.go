package sort

type QuickSort struct{}

func (s QuickSort) sort(A *[]int) {
	var qsort func(start, end int)

	var partition = func(start, end, pivot int) int {
		if start == end {
			return start
		}

		swap(A, pivot, start)

		left := start
		right := end

		for {

			for left < end && (*A)[start] >= (*A)[left] {
				left++
			}

			for right > start && (*A)[start] <= (*A)[right] {
				right--
			}

			if left >= right {
				break
			}

			swap(A, left, right)
		}

		swap(A, start, right)

		return right
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
