package itertools

import "iter"

func Cache[V any](i iter.Seq[V]) iter.Seq[V] {
	var c []V
	next, stop := iter.Pull(i)
	return func(yield func(V) bool) {
		for _,v  := range c {
			if !yield(v) {
				return
			}
		}
		for {
			v, ok := next()
			if !ok {
				stop()
				return
			}
			c = append(c, v)
			if !yield(v) {
				return
			}
		}
	}
}

func Cache2[K, V any](i iter.Seq2[K, V]) iter.Seq2[K, V] {
	return Make2[K, V](Cache)(i)
}
