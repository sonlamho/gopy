package gopy

import (
	"fmt"
	"testing"
)

func checkEq[T comparable, C ~[]T](got C, want C, t *testing.T) bool {
	if !Eq[T](got, want) {
		t.Errorf("Result = %s, want %s", fmt.Sprint(got), fmt.Sprint(want))
	}
	return true
}

func TestMap(t *testing.T) {

	t.Run("int slice", func(t *testing.T) {
		f := func(x int) int { return x + 1000 }
		slice := []int{5, 0, 10, 123, -1}
		want := []int{1005, 1000, 1010, 1123, 999}
		checkEq(Map(f, slice), want, t)
	})

	t.Run("empty slice", func(t *testing.T) {
		f := func(x int) int { return x + 1000 }
		slice := []int{}
		want := []int{}
		checkEq(Map(f, slice), want, t)
	})

	t.Run("uint16 slice", func(t *testing.T) {
		f := func(x uint16) uint16 { return x + 1000 }
		slice := []uint16{0, 1, 100, 999, 323}
		want := []uint16{1000, 1001, 1100, 1999, 1323}
		checkEq(Map(f, slice), want, t)
	})

	t.Run("float64 slice", func(t *testing.T) {
		f := func(x float64) float64 { return x + 1000.56 }
		slice := []float64{-0.56, 234.0, 1.0, 0.0}
		want := []float64{1000.0, 1234.56, 1001.56, 1000.56}
		checkEq(Map(f, slice), want, t)
	})

	t.Run("string slice", func(t *testing.T) {
		f := func(x string) string { return "-" + x + "-" }
		slice := []string{"", "0", "abcd", "", "hjkl"}
		want := []string{"--", "-0-", "-abcd-", "--", "-hjkl-"}
		checkEq(Map(f, slice), want, t)
	})
}

func TestFilter(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		pred := func(x int) bool { return true }
		slice := []int{}
		want := []int{}
		checkEq(Filter(pred, slice), want, t)
	})

	t.Run("always false predicate", func(t *testing.T) {
		pred := func(x int) bool { return false }
		slice := []int{1, 2, 3, 4, 5, 6}
		want := []int{}
		checkEq(Filter(pred, slice), want, t)
	})

	t.Run("always true predicate", func(t *testing.T) {
		pred := func(x uint32) bool { return true }
		slice := []uint32{1, 2, 3, 4, 5, 6}
		want := slice
		checkEq(Filter(pred, slice), want, t)
	})

	t.Run("filter even len strings", func(t *testing.T) {
		pred := func(x string) bool { return len(x)%2 == 0}
		slice := []string{"a", "", "qwer", "jkl", "jk", "12345"}
		want := []string{"", "qwer", "jk"}
		checkEq(Filter(pred, slice), want, t)
	})

	t.Run("filter high numbers", func(t *testing.T) {
		pred := func(x float64) bool { return x > 1000 }
		slice := []float64{32., 59., -23104., 12039., 1000.1, 999., 0., 9999.9}
		want := []float64{12039., 1000.1, 9999.9}
		checkEq(Filter(pred, slice), want, t)
	})

}
