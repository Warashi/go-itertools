package itertools

import "iter"

type Pair[A, B any] struct {
	A A
	B B
}

func PairSeq[A, B any](i iter.Seq2[A, B]) iter.Seq[Pair[A, B]] {
	return func(yield func(Pair[A, B]) bool) {
		for a, b := range i {
			if !yield(Pair[A, B]{a, b}) {
				break
			}
		}
	}
}

func UnpairSeq[A, B any](i iter.Seq[Pair[A, B]]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		for p := range i {
			if !yield(p.A, p.B) {
				break
			}
		}
	}
}

func PairFn[A, B, V any](f func(a A, b B) V) func(Pair[A, B]) V {
	return func(p Pair[A, B]) V {
		return f(p.A, p.B)
	}
}

func UnpairFn[A, B, V any](f func(Pair[A, B]) V) func(A, B) V {
	return func(a A, b B) V {
		return f(Pair[A, B]{a, b})
	}
}

func Make2[K, V any](f SeqModifier[Pair[K, V], Pair[K, V]]) Seq2Modifier[K, V, K, V] {
	return func(i iter.Seq2[K, V]) iter.Seq2[K, V] {
		return UnpairSeq(f(PairSeq(i)))
	}
}

func Make2Fn[K, V, R any](f func(func(Pair[K, V]) R) SeqModifier[Pair[K, V], Pair[K, V]]) func(func(K, V) R) Seq2Modifier[K, V, K, V] {
	return func(fn func(K, V) R) Seq2Modifier[K, V, K, V] {
		return Make2(f(PairFn(fn)))
	}
}
