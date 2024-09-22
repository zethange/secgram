package util

func Some[T any](array []T, compare func(T) bool) bool {
	for _, v := range array {
		if compare(v) {
			return true
		}
	}
	return false
}
