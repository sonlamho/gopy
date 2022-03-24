package gopy

// Number is a generic interface of all numeric types in Go.
type Number interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8 | float64 | float32
}

// NumLike is a generic interface of all numeric types and custom numeric types in Go.
type NumLike interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~float64 | ~float32
}

// Map takes a `function` and a `slice`, calls `function` on all elements of `slice` and returns the results in another slice.
func Map[T any, U any, C ~[]T](function func(T) U, slice C) []U {
	result := make([]U, len(slice))
	for i, x := range slice {
		result[i] = function(x)
	}
	return result
}

// Eq tests if two slice-like objects contain equal elements.
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

// Filter returns a subset of `slice` containing elements satisfying the predicate `pred`.
func Filter[T any, C ~[]T](pred func(T) bool, slice C) []T {
	result := make([]T, 0)
	for _, x := range slice {
		if pred(x) {
			result = append(result, x)
		}
	}
	return result
}

// Reduce does the reductiong `function` over the whole `sequence` and returns the final result.
func Reduce[T any, U any, Uslice ~[]U](function func(T, U) T, sequence Uslice, initial T) T {
	result := initial
	for _, x := range sequence {
		result = function(result, x)
	}
	return result
}

// Reversed returns a reversed copy of `sequence`.
func Reversed[T any, C ~[]T](sequence C) []T {
	result := make([]T, len(sequence))
	lastIdx := len(sequence) - 1
	for i, x := range sequence {
		result[lastIdx-i] = x
	}
	return result
}

// Sum returns the sum of the collection of numbers.
func Sum[T NumLike, NumSlice ~[]T](nums NumSlice) T {
	return Reduce(func(x, y T) T { return x + y }, nums, T(0))
}

// VarSum is the variadic version of Sum.
func VarSum[T Number](nums ...T) T {
	return Sum(nums)
}

func reduceMonoid[T any, Tslice ~[]T](function func(T, T) T, sequence Tslice, val0 T) T {
	if len(sequence) == 0 {
		return val0
	}
	return Reduce(function, sequence, sequence[0])

}

func min[T NumLike](x, y T) T {
	if x < y {
		return x
	}
	return y

}

// Min returns the minimum element of `nums`.
func Min[T NumLike, NumSlice ~[]T](nums NumSlice) T {
	return reduceMonoid(min[T], nums, T(0))
}

// VarMin is the variadic version of Min.
func VarMin[T Number](nums ...T) T {
	return Min(nums)
}

func max[T NumLike](x, y T) T {
	if x >= y {
		return x
	}
	return y

}

// Max returns the maximum element of `nums`.
func Max[T NumLike, NumSlice ~[]T](nums NumSlice) T {
	return reduceMonoid(max[T], nums, T(0))
}

// VarMax is the variadic version of Max.
func VarMax[T Number](nums ...T) T {
	return Max(nums)
}

// All returns true if and only if all elements of `bools` is true.
func All[BoolSlice ~[]bool](bools BoolSlice) bool {
	return Reduce(
		func(a, b bool) bool { return a && b },
		bools,
		true,
	)
}

// VarAll is the variadic version of All.
func VarAll(bools ...bool) bool {
	return All(bools)
}

// Any returns true if any element of `bools` is true, and returns false otherwise.
func Any[BoolSlice ~[]bool](bools BoolSlice) bool {
	return Reduce(
		func(a, b bool) bool { return a || b },
		bools,
		false,
	)
}

// VarAny is the variadic version of Any.
func VarAny(bools ...bool) bool {
	return Any(bools)
}
