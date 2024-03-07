package sort

type MergeSort struct{}

// Recursive, Top-Down Approach
func (s MergeSort) sort(A *[]int) {
	aux := make([]int, len(*A))

	var (
		merge func(lo int, mid int, hi int)
		rsort func(lo int, hi int)
	)

	merge = func(lo int, mid int, hi int) {
		copy(aux[lo:hi+1], (*A)[lo:hi+1])

		left := lo
		right := mid + 1

		for i := lo; i < hi+1; i++ {

			if left > mid {
				(*A)[i] = aux[right]
				right += 1
			} else if right > hi {
				(*A)[i] = aux[left]
				left += 1
			} else if aux[right] < aux[left] {
				(*A)[i] = aux[right]
				right += 1
			} else {
				(*A)[i] = aux[left]
				left += 1
			}
		}
	}

	rsort = func(lo int, hi int) {
		if hi <= lo {
			return
		}

		mid := lo + (hi-lo)/2

		rsort(lo, mid)
		rsort(mid+1, hi)

		merge(lo, mid, hi)
	}

	rsort(0, len(*A)-1)
}
