package xslice

// Diff 差集
func Diff[T comparable](sliceA, sliceB []T) (result []T) {
	set := make(map[T]struct{}, len(sliceB))

	for _, s := range sliceB {
		set[s] = struct{}{}
	}

	for _, s := range sliceA {
		if _, ok := set[s]; ok {
			delete(set, s) // remove duplicates
		} else {
			result = append(result, s)
		}
	}
	return result
}

// Intersect 交集
func Intersect[T comparable](ss []T, slices ...[]T) (ss2 []T) {
	if slices == nil {
		return nil
	}

	var unique = make([]map[T]struct{}, len(slices))
	for i := 0; i < len(slices); i++ {
		m := make(map[T]struct{})
		for _, el := range slices[i] {
			m[el] = struct{}{}
		}
		unique[i] = m
	}

	var containsInAll = false
	for _, el := range Unique(ss) {
		for _, u := range unique {
			if _, exists := u[el]; !exists {
				containsInAll = false
				break
			}
			containsInAll = true
		}
		if containsInAll {
			ss2 = append(ss2, el)
		}
	}

	return
}

// Unique 去重
func Unique[T comparable](ss []T) (result []T) {
	if len(ss) < 2 {
		return ss
	}

	tmp := map[T]struct{}{}

	for _, value := range ss {
		if _, ok := tmp[value]; !ok {
			tmp[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}
