package sort

type ShellSort struct{}

func (s ShellSort) sort(A *[]int) {
	n := len(*A)

	for h := n / 2; h > 0; h /= 2 {

		for i := h; i < n; i++ {
			tmp := (*A)[i]

			var j int

			for j = i; !less(j, h) && less(tmp, (*A)[j-h]); j -= h {
				(*A)[j] = (*A)[j-h]
			}

			(*A)[j] = tmp
		}
	}
}
