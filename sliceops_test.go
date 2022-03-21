package gopy

import (
	"fmt"
	"testing"
)

func checkEqSlice[T comparable, C ~[]T](got C, want C, t *testing.T) bool {
	if !Eq[T](got, want) {
		t.Errorf("Result = %s, want %s", fmt.Sprint(got), fmt.Sprint(want))
	}
	return true
}

func checkEq[T comparable](got T, want T, t *testing.T) bool {
	if got != want {
		t.Errorf("Result = %s, want %s", fmt.Sprint(got), fmt.Sprint(want))
	}
	return true
}

func TestMap(t *testing.T) {

	t.Run("int slice", func(t *testing.T) {
		f := func(x int) int { return x + 1000 }
		slice := []int{5, 0, 10, 123, -1}
		want := []int{1005, 1000, 1010, 1123, 999}
		checkEqSlice(Map(f, slice), want, t)
	})

	t.Run("empty slice", func(t *testing.T) {
		f := func(x int) int { return x + 1000 }
		slice := []int{}
		want := []int{}
		checkEqSlice(Map(f, slice), want, t)
	})

	t.Run("uint16 slice", func(t *testing.T) {
		f := func(x uint16) uint16 { return x + 1000 }
		slice := []uint16{0, 1, 100, 999, 323}
		want := []uint16{1000, 1001, 1100, 1999, 1323}
		checkEqSlice(Map(f, slice), want, t)
	})

	t.Run("float64 slice", func(t *testing.T) {
		f := func(x float64) float64 { return x + 1000.56 }
		slice := []float64{-0.56, 234.0, 1.0, 0.0}
		want := []float64{1000.0, 1234.56, 1001.56, 1000.56}
		checkEqSlice(Map(f, slice), want, t)
	})

	t.Run("string slice", func(t *testing.T) {
		f := func(x string) string { return "-" + x + "-" }
		slice := []string{"", "0", "abcd", "", "hjkl"}
		want := []string{"--", "-0-", "-abcd-", "--", "-hjkl-"}
		checkEqSlice(Map(f, slice), want, t)
	})
}

func TestFilter(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		pred := func(x int) bool { return true }
		slice := []int{}
		want := []int{}
		checkEqSlice(Filter(pred, slice), want, t)
	})

	t.Run("always false predicate", func(t *testing.T) {
		pred := func(x int) bool { return false }
		slice := []int{1, 2, 3, 4, 5, 6}
		want := []int{}
		checkEqSlice(Filter(pred, slice), want, t)
	})

	t.Run("always true predicate", func(t *testing.T) {
		pred := func(x uint32) bool { return true }
		slice := []uint32{1, 2, 3, 4, 5, 6}
		want := slice
		checkEqSlice(Filter(pred, slice), want, t)
	})

	t.Run("filter even len strings", func(t *testing.T) {
		pred := func(x string) bool { return len(x)%2 == 0 }
		slice := []string{"a", "", "qwer", "jkl", "jk", "12345"}
		want := []string{"", "qwer", "jk"}
		checkEqSlice(Filter(pred, slice), want, t)
	})

	t.Run("filter high numbers", func(t *testing.T) {
		pred := func(x float64) bool { return x > 1000 }
		slice := []float64{32., 59., -23104., 12039., 1000.1, 999., 0., 9999.9}
		want := []float64{12039., 1000.1, 9999.9}
		checkEqSlice(Filter(pred, slice), want, t)
	})

}

func TestReduce(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		fun := func(x int, y int) int { return x + y }
		seq := []int{}
		init := 1234
		want := init
		checkEq(Reduce(fun, seq, init), want, t)
	})

	t.Run("trivial function", func(t *testing.T) {
		fun := func(x float64, y int32) float64 { return 0.2 }
		seq := []int32{5, -123, 4000, 9999}
		init := 1234.56
		want := float64(0.2)
		checkEq(Reduce(fun, seq, init), want, t)
	})

	t.Run("summing function", func(t *testing.T) {
		fun := func(x uint32, y uint32) uint32 { return x + y }
		seq := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		init := uint32(0)
		want := uint32(6 * 13)
		checkEq(Reduce(fun, seq, init), want, t)
	})

	t.Run("cummulative sum", func(t *testing.T) {
		fun := func(x []int, y int) []int {
			last := 0
			if n := len(x); n > 0 {
				last = x[n-1]
			}
			return append(x, last+y)
		}
		seq := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3}
		init := make([]int, 0, len(seq))
		want := []int{10, 19, 27, 34, 40, 45, 49, 52, 54, 55, 55, 54, 52, 49}
		checkEqSlice(Reduce(fun, seq, init), want, t)
	})

	t.Run("reversing (inefficient)", func(t *testing.T) {
		fun := func(x []int, y int) []int {
			prefix := []int{y}
			return append(prefix, x...)
		}
		seq := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3}
		init := []int{}
		want := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		checkEqSlice(Reduce(fun, seq, init), want, t)
	})

}
