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

// Contains 检查 ss 是否包含 target
func Contains[T comparable](ss []T, target T) bool {
	for _, s := range ss {
		if s == target {
			return true
		}
	}
	return false
}

// Merge 合并 slice
func Merge[T any](slices ...[]T) []T {
	var newSlice []T
	for _, s := range slices {
		newSlice = append(newSlice, s...)
	}
	return newSlice
}

// FirstOr 获取第一个元素，如果不存在返回 elseVal
func FirstOr[T any](sl []T, elseVal T) T {
	if len(sl) > 0 {
		return sl[0]
	}
	return elseVal
}

// Reverse 反转
func Reverse[T any](ss []T) {
	ln := len(ss)
	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		ss[i], ss[li] = ss[li], ss[i]
	}
}

// IndexOf 获取 list 中第一个匹配  val 的下标，未匹配返回 -1
func IndexOf[T comparable](list []T, val T) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}
