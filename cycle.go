package itertools

import "iter"

func Cycle[V any](i iter.Seq[V]) iter.Seq[V] {
	var buf []V
	return func(yield func(V) bool) {
		for v := range i {
			buf = append(buf, v)
			if !yield(v) {
				return
			}
		}
		for {
			for _, v := range buf {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Cycle2[K, V any](i iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Make2[K, V](Cycle)(i)
}
