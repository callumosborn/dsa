package sort

type ShellSort struct{}

func (s ShellSort) sort(A *[]int) {
	n := len(*A)

	// h, is the gap between the values in the slice.
	// The gap is used to define the sub lists that are checked.
	for h := n / 2; h > 0; h /= 2 {

		for i := h; i < n; i++ {
			// tmp holds the current value that is being sorted.
			tmp := (*A)[i]

			var j int

			// !less(j, h) ensures that index j does not move past the beginning of the sub list.
			// less(tmp, (*A)[j-h]) ensures that the current value to be sorted is less that a previous value in the sub list.
			for j = i; !less(j, h) && less(tmp, (*A)[j-h]); j -= h {
				// Moves the greater element up by one interval gap in the sub list.
				(*A)[j] = (*A)[j-h]
			}

			// If the index j is before the start of the sub list or the current value to be sorted is not
			// less than the previous element in the sub list.
			// The current value to be sorted is put into it's place in the sub list.
			(*A)[j] = tmp
		}
	}
}
