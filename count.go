package itertools

import (
	"iter"

	"golang.org/x/exp/constraints"
)

func Count[V constraints.Integer | constraints.Float](start, step V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := V(start); ; i += step {
			if !yield(i) {
				break
			}
		}
	}
}
