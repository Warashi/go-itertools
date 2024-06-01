package itertools

import "iter"

func Filter[V any](fn func(V) bool) SeqModifier[V, V] {
	return func(i iter.Seq[V]) iter.Seq[V] {
		return func(yield func(V) bool) {
			for v := range i {
				if fn(v) {
					if !yield(v) {
						break
					}
				}
			}
		}
	}
}

func Filter2[K, V any](fn func(K, V) bool) Seq2Modifier[K, V, K, V] {
	return Make2(Filter(PairFn(fn)))
}
