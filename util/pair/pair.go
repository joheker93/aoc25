package pair

type Pair[K, V any] struct {
	Fst K
	Snd V
}

func Of[K, V any](k K, v V) Pair[K, V] {
	return Pair[K, V]{k, v}
}

func (pair Pair[K, V]) Unpack() (K, V) {
	return pair.Fst, pair.Snd
}
