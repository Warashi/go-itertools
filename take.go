package itertools

import (
	"iter"

	"golang.org/x/exp/constraints"
)

func Take[V any, N constraints.Integer](n N) SeqModifier[V, V] {
	return func(i iter.Seq[V]) iter.Seq[V] {
		return func(yield func(V) bool) {
			for v := range i {
				if n <= 0 {
					return
				}
				n--
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Take2[K, V any, N constraints.Integer](n N) Seq2Modifier[K, V, K, V] {
	return Make2(Take[Pair[K, V]](n))
}

func TakeWhile[V any](fn func(V) bool) SeqModifier[V, V] {
	return func(i iter.Seq[V]) iter.Seq[V] {
		return func(yield func(V) bool) {
			for v := range i {
				if !fn(v) {
					return
				}
				if !yield(v) {
					return
				}
			}
		}
	}
}

func TakeWhile2[K, V any](fn func(K, V) bool) Seq2Modifier[K, V, K, V] {
	return Make2Fn[K, V](TakeWhile)(fn)
}
