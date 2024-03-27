package sort

type MergeSort struct{}

// Recursive, Top-Down Approach
func (s MergeSort) sort(A *[]int) {
	aux := make([]int, len(*A))

	var (
		merge func(start, mid, end int)
		rsort func(start, end int)
	)

	merge = func(start, mid, end int) {
		copy(aux[start:end+1], (*A)[start:end+1])

		left := start
		right := mid + 1

		for i := start; i < end+1; i++ {

			if left > mid {
				(*A)[i] = aux[right]
				right += 1
			} else if right > end {
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

	rsort = func(start, end int) {
		if end == start {
			return
		}

		mid := start + (end-start)/2

		rsort(start, mid)
		rsort(mid+1, end)

		merge(start, mid, end)
	}

	rsort(0, len(*A)-1)
}

// Non-Recursive, Bottom-Up Approach
func (s MergeSort) sort2(A *[]int) {
	n := len(*A)
	aux := make([]int, n)

	var (
		min   func(a int, b int) int
		merge func(lo int, mid int, hi int)
	)

	min = func(a int, b int) int {
		if a < b {
			return a
		}

		return b
	}

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

	for i := 1; i <= n-1; i = i + i {

		for lo := 0; lo < n-i; lo += i + i {
			mid := min(lo+i-1, n-1)
			hi := min(lo+2*i-1, n-1)

			merge(lo, mid, hi)
		}
	}
}
