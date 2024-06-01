package itertools

import "iter"

func Chain[V any](mods ...SeqModifier[V, V]) SeqModifier[V, V] {
	return func(i iter.Seq[V]) iter.Seq[V] {
		return func(yield func(V) bool) {
			for _, mod := range mods {
				i = mod(i)
			}
			for v := range i {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func Chain2[K, V any](mods ...Seq2Modifier[K, V, K, V]) Seq2Modifier[K, V, K, V] {
	return func(i iter.Seq2[K, V]) iter.Seq2[K, V] {
		return func(yield func(K, V) bool) {
			for _, mod := range mods {
				i = mod(i)
			}
			for k, v := range i {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}
