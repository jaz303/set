package set

func AddMapKeys[K comparable, V any](dst Set[K], src map[K]V) {
	for k := range src {
		dst.Add(k)
	}
}

func AddMapValues[K comparable, V comparable](dst Set[V], src map[K]V) {
	for _, v := range src {
		dst.Add(v)
	}
}
