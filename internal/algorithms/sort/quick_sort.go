package sort

type QuickSort struct{}

func (s QuickSort) sort(A *[]int) {
	var qsort func(lo, hi int)

	var partition = func(lo, hi, pivot int) int {
		if lo == hi {
			return lo
		}

		swap(A, pivot, lo)

		i := lo
		j := hi

		for {

			for i < hi && (*A)[lo] >= (*A)[i] {
				i++
			}

			for j > lo && (*A)[lo] <= (*A)[j] {
				j--
			}

			if i >= j {
				break
			}

			swap(A, i, j)
		}

		swap(A, lo, j)

		return j
	}

	qsort = func(lo, hi int) {
		if hi <= lo {
			return
		}

		pivot_idx := lo
		spivot_idx := partition(lo, hi, pivot_idx)

		qsort(lo, spivot_idx-1)
		qsort(spivot_idx+1, hi)
	}

	qsort(0, len(*A)-1)
}
