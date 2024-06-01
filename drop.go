package itertools

import (
	"iter"

	"golang.org/x/exp/constraints"
)

func Drop[V any, N constraints.Integer](n N) SeqModifier[V, V] {
	return func(i iter.Seq[V]) iter.Seq[V] {
		return func(yield func(V) bool) {
			for v := range i {
				if n > 0 {
					n--
					continue
				}
				if !yield(v) {
					break
				}
			}
		}
	}
}

func Drop2[K, V any, N constraints.Integer](n N) Seq2Modifier[K, V, K, V] {
	return Make2(Drop[Pair[K, V]](n))
}

func DropWhile[V any](fn func(V) bool) SeqModifier[V, V] {
	return func(i iter.Seq[V]) iter.Seq[V] {
		return func(yield func(V) bool) {
			dropping := true
			for v := range i {
				if dropping {
					if !fn(v) {
						dropping = false
					}
					continue
				}
				if !yield(v) {
					break
				}
			}
		}
	}
}

func DropWhile2[K, V any](fn func(K, V) bool) Seq2Modifier[K, V, K, V] {
	return Make2Fn[K, V](DropWhile)(fn)
}
