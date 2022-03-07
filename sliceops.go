package gopy

type Number interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8 | float64 | float32
}

func Map[T any, U any, C ~[]T](f func(T) U, slice C) []U {
	result := make([]U, len(slice))
	for i, x := range slice {
		result[i] = f(x)
	}
	return result
}

func Eq[T comparable, C ~[]T](a, b C) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Filter[T any, C ~[]T](pred func(T) bool, slice C) []T {
	result := make([]T, 0)
	for _, x := range slice {
		if pred(x) {
			result = append(result, x)
		}
	}
	return result
}
