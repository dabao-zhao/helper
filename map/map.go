package xmap

// Keys 获取 map 的所有 key
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

// Values 获取 map 的所有 value
func Values[K comparable, V any](m map[K]V) []V {
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]V, l)
	for _, value := range m {
		keys[i] = value
		i++
	}

	return keys
}
