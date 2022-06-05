package types

import (
	"testing"

	u "github.com/Benschar1/golonzo/utils"
)

func TestIsSome(t *testing.T) {
	cases := map[string]u.FuncTc{
		"Some{x} is some": u.MakeFtc1(
			IsSome[int],
			true,
			Some[int]{4},
		),
		"None{} is not some": u.MakeFtc1(
			IsSome[int],
			false,
			None[int]{},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestIsNone(t *testing.T) {
	cases := map[string]u.FuncTc{
		"Some{x} is not none": u.MakeFtc1(
			IsNone[int],
			false,
			Some[int]{4},
		),
		"None{} is none": u.MakeFtc1(
			IsNone[int],
			true,
			None[int]{},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestFromMaybe(t *testing.T) {
	cases := map[string]u.FuncTc{
		"FromMaybe on Some{x} is x": u.MakeFtc1(
			FromMaybe[int],
			3,
			4, Some[int]{3},
		),
		"FromMaybe on None{} is default": u.MakeFtc1(
			FromMaybe[int],
			4,
			4, None[int]{},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestMapMaybe(t *testing.T) {
	plus2 := func(i int) int { return i + 2 }

	cases := map[string]u.FuncTc{
		"MapMaybe on Some{x} is mapped x": u.MakeFtc1(
			MapMaybe[int, int],
			Some[int]{5},
			plus2, Some[int]{3},
		),
		"MapMaybe on None{} is none": u.MakeFtc1(
			MapMaybe[int, int],
			None[int]{},
			plus2, None[int]{},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestBindMaybe(t *testing.T) {
	maybePos := func(n int) Maybe[int] {
		if n > 0 {
			return Some[int]{n}
		}
		return None[int]{}
	}

	cases := map[string]u.FuncTc{
		"BindMaybe 1": u.MakeFtc1(
			BindMaybe[int, int],
			Some[int]{3},
			Some[int]{3}, maybePos,
		),
		"BindMaybe 2": u.MakeFtc1(
			BindMaybe[int, int],
			None[int]{},
			Some[int]{-1}, maybePos,
		),
		"BindMaybe 3": u.MakeFtc1(
			BindMaybe[int, int],
			None[int]{},
			None[int]{}, maybePos,
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestFilterMaybes(t *testing.T) {
	cases := map[string]u.FuncTc{
		"FilterMaybes filters all nones": u.MakeFtc1(
			FilterMaybes[int],
			[]int{1, -4, 5, -10},
			[]Maybe[int]{
				Some[int]{1},
				Some[int]{-4},
				None[int]{},
				Some[int]{5},
				None[int]{},
				Some[int]{-10},
			},
		),
		"FilterMaybes on all nones returns empty list": u.MakeFtc1(
			FilterMaybes[int],
			[]int{},
			[]Maybe[int]{
				None[int]{},
				None[int]{},
				None[int]{},
				None[int]{},
			},
		),
		"FilterMaybes on all somes returns whole list": u.MakeFtc1(
			FilterMaybes[int],
			[]int{1, 4, 7, -12, -993},
			[]Maybe[int]{
				Some[int]{1},
				Some[int]{4},
				Some[int]{7},
				Some[int]{-12},
				Some[int]{-993},
			},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}
