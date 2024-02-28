package sort

import "errors"

type Sorter interface {
	sort(A *[]int)
}

func less(a int, b int) bool {
	return a < b
}

func swap(A *[]int, a int, b int) error {
	if len(*A) <= 1 {
		return errors.New("invalid slice: slices cannot have less than 2 elements")
	}

	if a < 0 || b < 0 {
		return errors.New("invalid index: indices cannot be negative")
	}

	if a > len(*A)-1 || b > len(*A)-1 {
		return errors.New("invalid index: indices cannot be out-of-range")
	}

	(*A)[a], (*A)[b] = (*A)[b], (*A)[a]

	return nil
}
