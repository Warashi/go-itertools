package itertools

import "iter"

func ForAll[V any](fn func (V) bool) func(iter.Seq[V]) bool {
	return func(i iter.Seq[V]) bool {
		for v := range i {
			if !fn(v) {
				return false
			}
		}
		return true
	}
}
