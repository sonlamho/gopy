package gopy

import (
	"fmt"
	"math/rand"
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

type myInt int
type myFloat float64
type myStr string

func TestMap(t *testing.T) {

	t.Run("int slice", func(t *testing.T) {
		f := func(x int) int { return x + 1000 }
		slice := []int{5, 0, 10, 123, -1}
		want := []int{1005, 1000, 1010, 1123, 999}
		checkEqSlice(Map(f, slice), want, t)
	})

	t.Run("myInt slice", func(t *testing.T) {
		f := func(x myInt) myInt { return x + 1000 }
		slice := []myInt{5, 0, 10, 123, -1}
		want := []myInt{1005, 1000, 1010, 1123, 999}
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

	t.Run("filter even len myStr's", func(t *testing.T) {
		pred := func(x myStr) bool { return len(x)%2 == 0 }
		slice := []myStr{"a", "qwer", "jkl;", "jk", "12345"}
		want := []myStr{"qwer", "jkl;", "jk"}
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
		fun := func(x []myInt, y myInt) []myInt {
			last := myInt(0)
			if n := len(x); n > 0 {
				last = x[n-1]
			}
			return append(x, last+y)
		}
		seq := []myInt{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3}
		init := make([]myInt, 0, len(seq))
		want := []myInt{10, 19, 27, 34, 40, 45, 49, 52, 54, 55, 55, 54, 52, 49}
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

func TestReversed(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		seq := []int32{}
		want := []int32{}
		checkEqSlice(Reversed(seq), want, t)
	})

	t.Run("1 element slice", func(t *testing.T) {
		x := rand.Float64()
		seq := []float64{x}
		want := []float64{x}
		checkEqSlice(Reversed(seq), want, t)
	})

	t.Run("myInt slice", func(t *testing.T) {
		seq := []myInt{1, 2, 3, 4, 5, 6, 7}
		want := []myInt{7, 6, 5, 4, 3, 2, 1}
		checkEqSlice(Reversed(seq), want, t)
	})

	t.Run("Reversed vs implementation using Reduce", func(t *testing.T) {
		fun := func(x []int, y int) []int {
			prefix := []int{y}
			return append(prefix, x...)
		}
		seq := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3}
		result0 := Reversed(seq)
		result1 := Reduce(fun, seq, []int{})
		checkEqSlice(result0, result1, t)
	})

}

func TestSum(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		seq := []int{}
		want := 0
		checkEq(Sum(seq), want, t)
		checkEq(VarSum(seq...), want, t)
	})

	t.Run("uint32 slice", func(t *testing.T) {
		seq := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		want := uint32(6 * 13)
		checkEq(Sum(seq), want, t)
		checkEq(VarSum(seq...), want, t)
	})

	t.Run("int slice", func(t *testing.T) {
		seq := []int{-5, -4, -3, -2, -1, 0, 1234567, 0, 5, 4, 3, 2, 1}
		want := int(1234567)
		checkEq(Sum(seq), want, t)
		checkEq(VarSum(seq...), want, t)
	})

	t.Run("myInt slice", func(t *testing.T) {
		seq := []myInt{-5, -4, -3, -2, -1, 0, 1234567, 0, 5, 4, 3, 2, 1}
		want := myInt(1234567)
		checkEq(Sum(seq), want, t)
		// checkEq(VarSum(seq...), want, t)
	})

	t.Run("float64 slice", func(t *testing.T) {
		seq := []float64{1.5, -0.25, 1.0, 0.0, 0.75, 5.25}
		want := float64(8.25)
		checkEq(Sum(seq), want, t)
		checkEq(VarSum(seq...), want, t)
	})

}

func TestMin(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		seq := []int{}
		want := 0
		checkEq(Min(seq), want, t)
		checkEq(VarMin(seq...), want, t)
	})

	t.Run("uint32 slice", func(t *testing.T) {
		seq := []uint32{100, 50, 9, 423, 10, 12}
		want := uint32(9)
		checkEq(Min(seq), want, t)
		checkEq(VarMin(seq...), want, t)
	})

	t.Run("float64 slice", func(t *testing.T) {
		seq := []float64{1.5, -0.25, 1.0, 0.0, 0.75, 5.25, -0.2499}
		want := float64(-0.25)
		checkEq(Min(seq), want, t)
		checkEq(VarMin(seq...), want, t)
	})

	t.Run("myFloat slice", func(t *testing.T) {
		seq := []myFloat{1.5, -0.25, 1.0, 0.0, 0.75, 5.25, -0.2499}
		want := myFloat(-0.25)
		checkEq(Min(seq), want, t)
		// checkEq(VarMin(seq...), want, t)
	})

}

func TestMax(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		seq := []int{}
		want := 0
		checkEq(Max(seq), want, t)
		checkEq(VarMax(seq...), want, t)
	})

	t.Run("uint32 slice", func(t *testing.T) {
		seq := []uint32{100, 50, 9, 423, 10, 12}
		want := uint32(423)
		checkEq(Max(seq), want, t)
		checkEq(VarMax(seq...), want, t)
	})

	t.Run("myInt slice", func(t *testing.T) {
		seq := []myInt{100, 50, 9, 423, 10, 12}
		want := myInt(423)
		checkEq(Max(seq), want, t)
		// checkEq(VarMax(seq...), want, t)
	})

	t.Run("float64 slice", func(t *testing.T) {
		seq := []float64{1.5, -0.25, 1.0, 0.0, 0.75, 5.25, -0.2499}
		want := float64(5.25)
		checkEq(Max(seq), want, t)
		checkEq(VarMax(seq...), want, t)
	})

}

func TestAll(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		seq := []bool{}
		want := true
		checkEq(All(seq), want, t)
		checkEq(VarAll(seq...), want, t)
	})

	t.Run("all is true", func(t *testing.T) {
		seq := []bool{true, true, true, true}
		want := true
		checkEq(All(seq), want, t)
		checkEq(VarAll(seq...), want, t)
	})

	t.Run("not all 1", func(t *testing.T) {
		seq := []bool{true, true, true, false, true}
		want := false
		checkEq(All(seq), want, t)
		checkEq(VarAll(seq...), want, t)
	})

	t.Run("not all 2", func(t *testing.T) {
		seq := []bool{false, true, true, true, false, true}
		want := false
		checkEq(All(seq), want, t)
		checkEq(VarAll(seq...), want, t)
	})

}

func TestAny(t *testing.T) {

	t.Run("empty slice", func(t *testing.T) {
		seq := []bool{}
		want := false
		checkEq(Any(seq), want, t)
		checkEq(VarAny(seq...), want, t)
	})

	t.Run("any is true 1", func(t *testing.T) {
		seq := []bool{false, false, false, true, false}
		want := true
		checkEq(Any(seq), want, t)
		checkEq(VarAny(seq...), want, t)
	})

	t.Run("any is true 2", func(t *testing.T) {
		seq := []bool{true, false, false, true, false}
		want := true
		checkEq(Any(seq), want, t)
		checkEq(VarAny(seq...), want, t)
	})

	t.Run("not any", func(t *testing.T) {
		seq := []bool{false, false, false, false, false, false}
		want := false
		checkEq(Any(seq), want, t)
		checkEq(VarAny(seq...), want, t)
	})

}

func TestVariadic(t *testing.T) {

	t.Run("no args", func(t *testing.T) {
		checkEq(VarAny(), false, t)
		checkEq(VarAll(), true, t)
		checkEq(VarSum[float64](), 0.0, t)
		checkEq(VarMin[uint32](), uint32(0), t)
		checkEq(VarMax[int64](), int64(0), t)
	})

	t.Run("one arg All Any", func(t *testing.T) {
		for _, b := range []bool{false, true} {
			checkEq(VarAll(b), b, t)
			checkEq(VarAny(b), b, t)
		}

	})

	t.Run("one arg Sum Min Max", func(t *testing.T) {
		x := rand.Float64()
		checkEq(VarSum(x), x, t)
		checkEq(VarMin(x), x, t)
		checkEq(VarMax(x), x, t)

		y := rand.Int()
		checkEq(VarSum(y), y, t)
		checkEq(VarMin(y), y, t)
		checkEq(VarMax(y), y, t)
	})

	t.Run("two args", func(t *testing.T) {
		checkEq(VarAll(true, false), false, t)
		checkEq(VarAll(true, true), true, t)

		checkEq(VarAny(false, true), true, t)
		checkEq(VarAny(false, false), false, t)

		checkEq(VarMax(4, 3), 4, t)
		checkEq(VarMin(4, 3), 3, t)
		checkEq(VarSum(4, 3), 7, t)

		checkEq(VarMax(-4.75, -4.25), -4.25, t)
		checkEq(VarMin(-4.75, -4.25), -4.75, t)
		checkEq(VarSum(-4.75, -4.25), -9.0, t)

	})

}
