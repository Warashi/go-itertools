package itertools

import "iter"

type (
	SeqModifier[V1, V2 any]          func(iter.Seq[V1]) iter.Seq[V2]
	Seq2Modifier[K1, V1, K2, V2 any] func(iter.Seq2[K1, V1]) iter.Seq2[K2, V2]
)
