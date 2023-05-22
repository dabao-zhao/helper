package _map

func Keys[K comparable, V any](m map[K]V) []K {
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]K, l)
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
