package slicetl

import "iter"

func GetSlidingWindow[T any](field []T, windowSize int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for i := 0; i+windowSize <= len(field); i += windowSize {
			if !yield(field[i : i+windowSize]) {
				return
			}
		}
	}
}
