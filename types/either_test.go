package types

import (
	"testing"

	u "github.com/Benschar1/golonzo/utils"
)

func TestIsLeft(t *testing.T) {
	cases := map[string]u.FuncTc{
		"Left{...} is left": u.MakeFtc1(
			IsLeft[int, string],
			true,
			Left[int, string]{2},
		),
		"Right{...} isn't left": u.MakeFtc1(
			IsLeft[int, string],
			false,
			Right[int, string]{"st"},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestIsRight(t *testing.T) {
	cases := map[string]u.FuncTc{
		"Left{...} isn't right": u.MakeFtc1(
			IsRight[int, string],
			false,
			Left[int, string]{2},
		),
		"Right{...} is right": u.MakeFtc1(
			IsRight[int, string],
			true,
			Right[int, string]{"st"},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestFromLeft(t *testing.T) {
	cases := map[string]u.FuncTc{
		"FromLeft on Left{x} gives x": u.MakeFtc1(
			FromLeft[int, int],
			3,
			0, Left[int, int]{3},
		),
		"FromLeft on Right{x} gives default value": u.MakeFtc1(
			FromLeft[int, int],
			0,
			0, Right[int, int]{3},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestFromRight(t *testing.T) {
	cases := map[string]u.FuncTc{
		"FromRight on Left{x} gives default value": u.MakeFtc1(
			FromRight[int, int],
			0,
			0, Left[int, int]{3},
		),
		"FromRight on Right{x} gives x": u.MakeFtc1(
			FromRight[int, int],
			3,
			0, Right[int, int]{3},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestEitherMapL(t *testing.T) {
	incr := func(n int) int { return n + 1 }

	cases := map[string]u.FuncTc{
		"EitherMapL on Left{...} gives mapped value": u.MakeFtc1(
			EitherMapL[int, int, int],
			Left[int, int]{2},
			incr, Left[int, int]{1},
		),
		"EitherMapL on Right{...} gives default value": u.MakeFtc1(
			EitherMapL[int, int, int],
			Right[int, int]{1},
			incr, Right[int, int]{1},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestEitherMapR(t *testing.T) {
	incr := func(n int) int { return n + 1 }

	cases := map[string]u.FuncTc{
		"EitherMapR on Left{...} gives default value": u.MakeFtc1(
			EitherMapR[int, int, int],
			Left[int, int]{1},
			incr, Left[int, int]{1},
		),
		"EitherMapR on Right{...} gives mapped value": u.MakeFtc1(
			EitherMapR[int, int, int],
			Right[int, int]{2},
			incr, Right[int, int]{1},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestEitherMapLR(t *testing.T) {
	incr := func(n int) int { return n + 1 }
	decr := func(n int) int { return n - 1 }

	cases := map[string]u.FuncTc{
		"EitherMapLR on Left{...} maps with left function": u.MakeFtc1(
			EitherMapLR[int, int, int, int],
			Left[int, int]{2},
			incr, decr, Left[int, int]{1},
		),
		"EitherMapLR on Right{...} maps with right function": u.MakeFtc1(
			EitherMapLR[int, int, int, int],
			Right[int, int]{0},
			incr, decr, Right[int, int]{1},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestLeftToMaybe(t *testing.T) {
	cases := map[string]u.FuncTc{
		"LeftToMaybe on Left{x} returns Some{x}": u.MakeFtc1(
			LeftToMaybe[int, int],
			Some[int]{1},
			Left[int, int]{1},
		),
		"LeftToMaybe on Right{...} returns None": u.MakeFtc1(
			LeftToMaybe[int, int],
			None[int]{},
			Right[int, int]{1},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestRightToMaybe(t *testing.T) {
	cases := map[string]u.FuncTc{
		"RightToMaybe on Left{x} returns None": u.MakeFtc1(
			RightToMaybe[int, int],
			None[int]{},
			Left[int, int]{1},
		),
		"RightToMaybe on Right{...} returns Some{x}": u.MakeFtc1(
			RightToMaybe[int, int],
			Some[int]{1},
			Right[int, int]{1},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestUnifyEither(t *testing.T) {
	cases := map[string]u.FuncTc{
		"UnifyEither on Left{x} returns x": u.MakeFtc1(
			UnifyEither[int],
			3,
			Left[int, int]{3},
		),
		"UnifyEither on Right{x} returns x": u.MakeFtc1(
			UnifyEither[int],
			3,
			Right[int, int]{3},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestEitherFlip(t *testing.T) {
	cases := map[string]u.FuncTc{
		"EitherFlip on Left{x} returns Right{x}": u.MakeFtc1(
			EitherFlip[int, string],
			Right[string, int]{3},
			Left[int, string]{3},
		),
		"EitherFlip on Right{x} returns Left{x}": u.MakeFtc1(
			EitherFlip[string, int],
			Left[int, string]{3},
			Right[string, int]{3},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestLefts(t *testing.T) {
	cases := map[string]u.FuncTc{
		"Lefts on []Either returns only lefts": u.MakeFtc1(
			Lefts[int, string],
			[]int{3, 4, 0, -912, 3099},
			[]Either[int, string]{
				Left[int, string]{3},
				Right[int, string]{"daniuh"},
				Left[int, string]{4},
				Right[int, string]{"98u9"},
				Right[int, string]{""},
				Left[int, string]{0},
				Left[int, string]{-912},
				Right[int, string]{"synactic sugar causes cancer of the semicolon"},
				Left[int, string]{3099},
				Right[int, string]{"when will java die out already"},
				Right[int, string]{"public static void main(String[] args)"},
			},
		),
		"Lefts on []Either with all lefts returns whole list": u.MakeFtc1(
			Lefts[float64, float64],
			[]float64{3, -3.9, 3.1415},
			[]Either[float64, float64]{
				Left[float64, float64]{3.0},
				Left[float64, float64]{-3.9},
				Left[float64, float64]{3.1415},
			},
		),
		"Lefts on []Either with all rights returns empty list": u.MakeFtc1(
			Lefts[float64, float64],
			[]float64{},
			[]Either[float64, float64]{
				Right[float64, float64]{3.0},
				Right[float64, float64]{-3.9},
				Right[float64, float64]{3.1415},
			},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestRights(t *testing.T) {
	cases := map[string]u.FuncTc{
		"Rights on []Either returns only lefts": u.MakeFtc1(
			Rights[int, string],
			[]string{
				"daniuh",
				"98u9",
				"",
				"synactic sugar causes cancer of the semicolon",
				"when will java die out already",
				"public static void main(String[] args)",
			},
			[]Either[int, string]{
				Left[int, string]{3},
				Right[int, string]{"daniuh"},
				Left[int, string]{4},
				Right[int, string]{"98u9"},
				Right[int, string]{""},
				Left[int, string]{0},
				Left[int, string]{-912},
				Right[int, string]{"synactic sugar causes cancer of the semicolon"},
				Left[int, string]{3099},
				Right[int, string]{"when will java die out already"},
				Right[int, string]{"public static void main(String[] args)"},
			},
		),
		"Rights on []Either with all lefts returns empty list": u.MakeFtc1(
			Rights[float64, float64],
			[]float64{},
			[]Either[float64, float64]{
				Left[float64, float64]{3.0},
				Left[float64, float64]{-3.9},
				Left[float64, float64]{3.1415},
			},
		),
		"Rights on []Either with all rights returns whole list": u.MakeFtc1(
			Rights[float64, float64],
			[]float64{3, -3.9, 3.1415},
			[]Either[float64, float64]{
				Right[float64, float64]{3.0},
				Right[float64, float64]{-3.9},
				Right[float64, float64]{3.1415},
			},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}

func TestPartitionEither(t *testing.T) {

	cases := map[string]u.FuncTc{
		"PartitionEither correctly partitions list into lefts and rights": u.MakeFtc1(
			PartitionEither[int, string],
			Tuple2[[]int, []string]{
				[]int{3, 4, 0, -912, 3099},
				[]string{
					"daniuh",
					"98u9",
					"",
					"synactic sugar causes cancer of the semicolon",
					"when will java die out already",
					"public static void main(String[] args)",
				},
			},
			[]Either[int, string]{
				Left[int, string]{3},
				Right[int, string]{"daniuh"},
				Left[int, string]{4},
				Right[int, string]{"98u9"},
				Right[int, string]{""},
				Left[int, string]{0},
				Left[int, string]{-912},
				Right[int, string]{"synactic sugar causes cancer of the semicolon"},
				Left[int, string]{3099},
				Right[int, string]{"when will java die out already"},
				Right[int, string]{"public static void main(String[] args)"},
			},
		),
		"PartitionEither on all lefts returns partition with all lefts": u.MakeFtc1(
			PartitionEither[float64, float64],
			Tuple2[[]float64, []float64]{
				[]float64{3, -3.9, 3.1415},
				[]float64{},
			},
			[]Either[float64, float64]{
				Left[float64, float64]{3.0},
				Left[float64, float64]{-3.9},
				Left[float64, float64]{3.1415},
			},
		),
		"PartitionEither on all rights returns partition with all rights": u.MakeFtc1(
			PartitionEither[float64, float64],
			Tuple2[[]float64, []float64]{
				[]float64{},
				[]float64{3, -3.9, 3.1415},
			},
			[]Either[float64, float64]{
				Right[float64, float64]{3.0},
				Right[float64, float64]{-3.9},
				Right[float64, float64]{3.1415},
			},
		),
	}

	u.FuncUnitTests{Cases: cases}.Execute(t)
}
