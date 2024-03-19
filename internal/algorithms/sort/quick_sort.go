package sort

type QuickSort struct{}

func (s QuickSort) sort(A *[]int) {
	var qsort func(lo, hi int)

	var partition = func(A *[]int, lo int, hi int, pivot int) int {
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

		pivot := lo
		location := partition(A, lo, hi, pivot)

		qsort(lo, location-1)
		qsort(location+1, hi)
	}

	qsort(0, len(*A)-1)
}
