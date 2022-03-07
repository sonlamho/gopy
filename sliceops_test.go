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
