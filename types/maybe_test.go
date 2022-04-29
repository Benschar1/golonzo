package types

import (
	"fmt"
	"testing"

	"github.com/Benschar1/golonzo/utils"
)

func TestIsSome(t *testing.T) {
	test := utils.AsrtCallEq(t, "IsSome")

	in1, ex1 := Some[int]{4}, true
	in2, ex2 := None[int]{}, false

	test(ex1, IsSome[int](in1), in1)
	test(ex2, IsSome[int](in2), in2)
}

func TestIsNone(t *testing.T) {
	test := utils.AsrtCallEq(t, "IsNone")

	in1, ex1 := Some[int]{4}, false
	in2, ex2 := None[int]{}, true

	test(ex1, IsNone[int](in1), in1)
	test(ex2, IsNone[int](in2), in2)
}

func TestFromMaybe(t *testing.T) {

	test := utils.AsrtCallEq(t, "FromMaybe")

	df1, m1, ex1 := 4, Some[int]{3}, 3
	df2, m2, ex2 := 4, None[int]{}, 4

	test(ex1, FromMaybe(df1)(m1), df1, m1)
	test(ex2, FromMaybe(df2)(m2), df2, m2)
}

func TestMapMaybe(t *testing.T) {

	test := utils.AsrtCallEq(t, "MapMaybe")

	f1, in1, ex1 := func(n int) int { return n + 2 }, Some[int]{3}, Some[int]{5}
	f2, in2, ex2 := func(a int) string { return fmt.Sprint(a) }, None[int]{}, None[string]{}

	test(ex1, MapMaybe(f1)(in1), f1, in1)
	test(ex2, MapMaybe(f2)(in2), f2, in2)
}

func TestBindMaybe(t *testing.T) {

	test := utils.AsrtCallEq(t, "BindMaybe")

	maybePos := func(n int) Maybe[int] {
		if n > 0 {
			return Some[int]{n}
		}
		return None[int]{}
	}

	in1, ex1 := Some[int]{3}, Some[int]{3}
	in2, ex2 := Some[int]{-1}, None[int]{}
	in3, ex3 := None[int]{}, None[int]{}

	test(ex1, BindMaybe[int, int](in1)(maybePos), in1, maybePos)
	test(ex2, BindMaybe[int, int](in2)(maybePos), in2, maybePos)
	test(ex3, BindMaybe[int, int](in3)(maybePos), in3, maybePos)
}

func TestFilterMaybes(t *testing.T) {

	test := utils.AsrtCallEq(t, "FilterMaybes")

	l1, ex1 := []Maybe[int]{
		Some[int]{1},
		Some[int]{-4},
		None[int]{},
		Some[int]{5},
		None[int]{},
		Some[int]{-10},
	},
		[]int{1, -4, 5, -10}

	l2, ex2 := []Maybe[int]{
		None[int]{},
		None[int]{},
		None[int]{},
		None[int]{},
	},
		[]int{}

	l3, ex3 := []Maybe[int]{
		Some[int]{1},
		Some[int]{4},
		Some[int]{7},
		Some[int]{-12},
		Some[int]{-993},
	},
		[]int{1, 4, 7, -12, -993}

	test(ex1, FilterMaybes(l1), l1)
	test(ex2, FilterMaybes(l2), l2)
	test(ex3, FilterMaybes(l3), l3)
}
