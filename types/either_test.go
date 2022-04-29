package types

import (
	"testing"

	"github.com/Benschar1/golonzo/utils"
)

func TestIsLeft(t *testing.T) {
	test := utils.AsrtCallEq(t, "IsLeft")

	in1, ex1 := Left[int, string]{2}, true
	in2, ex2 := Right[int, string]{"st"}, false

	test(ex1, IsLeft[int, string](in1), in1)
	test(ex2, IsLeft[int, string](in2), in2)
}

func TestIsRight(t *testing.T) {
	test := utils.AsrtCallEq(t, "IsRight")

	in1, ex1 := Left[int, string]{2}, false
	in2, ex2 := Right[int, string]{"st"}, true

	test(ex1, IsRight[int, string](in1), in1)
	test(ex2, IsRight[int, string](in2), in2)
}

func TestFromLeft(t *testing.T) {
	test := utils.AsrtCallEq(t, "FromLeft")

	in1a, in1b, ex1 := 0, Left[int, int]{3}, 3
	in2a, in2b, ex2 := 0, Right[int, int]{3}, 0

	test(ex1, FromLeft[int, int](in1a)(in1b), in1a, in1b)
	test(ex2, FromLeft[int, int](in2a)(in2b), in2a, in2b)
}

func TestFromRight(t *testing.T) {
	test := utils.AsrtCallEq(t, "FromRight")

	in1a, in1b, ex1 := 0, Left[int, int]{3}, 0
	in2a, in2b, ex2 := 0, Right[int, int]{3}, 3

	test(ex1, FromRight[int](in1a)(in1b), in1a, in1b)
	test(ex2, FromRight[int](in2a)(in2b), in2a, in2b)
}

func TestEitherMapL(t *testing.T) {
	test := utils.AsrtCallEq(t, "EitherMapL")

	incr := func(n int) int { return n + 1 }

	in1, ex1 := Left[int, int]{1}, Left[int, int]{2}
	in2, ex2 := Right[int, int]{1}, Right[int, int]{1}

	test(ex1, EitherMapL[int, int, int](incr)(in1), incr, in1)
	test(ex2, EitherMapL[int, int, int](incr)(in2), incr, in2)
}

func TestEitherMapR(t *testing.T) {
	test := utils.AsrtCallEq(t, "EitherMapR")

	incr := func(n int) int { return n + 1 }

	in1, ex1 := Left[int, int]{1}, Left[int, int]{1}
	in2, ex2 := Right[int, int]{1}, Right[int, int]{2}

	test(ex1, EitherMapR[int, int, int](incr)(in1), incr, in1)
	test(ex2, EitherMapR[int, int, int](incr)(in2), incr, in2)
}

func TestEitherMapLR(t *testing.T) {
	test := utils.AsrtCallEq(t, "EitherMapLR")

	incr := func(n int) int { return n + 1 }
	decr := func(n int) int { return n - 1 }

	in1, ex1 := Left[int, int]{5}, Left[int, int]{6}
	in2, ex2 := Right[int, int]{5}, Right[int, int]{4}

	test(ex1, EitherMapLR[int, int, int, int](incr)(decr)(in1), incr, decr, in1)
	test(ex2, EitherMapLR[int, int, int, int](incr)(decr)(in2), incr, decr, in2)
}

func TestLeftToMaybe(t *testing.T) {
	test := utils.AsrtCallEq(t, "LeftToMaybe")

	in1, ex1 := Left[int, int]{3}, Some[int]{3}
	in2, ex2 := Right[int, int]{3}, None[int]{}

	test(ex1, LeftToMaybe[int, int](in1), in1)
	test(ex2, LeftToMaybe[int, int](in2), in2)
}

func TestRightToMaybe(t *testing.T) {
	test := utils.AsrtCallEq(t, "RightToMaybe")

	in1, ex1 := Left[int, int]{3}, None[int]{}
	in2, ex2 := Right[int, int]{3}, Some[int]{3}

	test(ex1, RightToMaybe[int, int](in1), in1)
	test(ex2, RightToMaybe[int, int](in2), in2)
}

func TestUnifyEither(t *testing.T) {
	test := utils.AsrtCallEq(t, "UnifyEither")

	in1, ex1 := Left[int, int]{3}, 3
	in2, ex2 := Right[int, int]{3}, 3

	test(ex1, UnifyEither[int](in1), in1)
	test(ex2, UnifyEither[int](in2), in2)
}

func TestEitherFlip(t *testing.T) {
	test := utils.AsrtCallEq(t, "EitherFlip")

	in1, ex1 := Left[int, int]{3}, Right[int, int]{3}
	in2, ex2 := Right[int, int]{3}, Left[int, int]{3}

	test(ex1, EitherFlip[int, int](in1), in1)
	test(ex2, EitherFlip[int, int](in2), in2)
}

func TestLefts(t *testing.T) {
	test := utils.AsrtCallEq(t, "Lefts")

	in1, ex1 := []Either[int, string]{
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
		[]int{3, 4, 0, -912, 3099}

	in2, ex2 := []Either[float64, float64]{
		Left[float64, float64]{3.0},
		Left[float64, float64]{-3.9},
		Left[float64, float64]{3.1415},
	},
		[]float64{3, -3.9, 3.1415}

	in3, ex3 := []Either[float64, float64]{
		Right[float64, float64]{3.0},
		Right[float64, float64]{-3.9},
		Right[float64, float64]{3.1415},
	},
		[]float64{}

	test(ex1, Lefts(in1), in1)
	test(ex2, Lefts(in2), in2)
	test(ex3, Lefts(in3), in3)
}

func TestRights(t *testing.T) {
	test := utils.AsrtCallEq(t, "Rights")

	in1, ex1 := []Either[int, string]{
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
		[]string{
			"daniuh",
			"98u9",
			"",
			"synactic sugar causes cancer of the semicolon",
			"when will java die out already",
			"public static void main(String[] args)",
		}

	in2, ex2 := []Either[float64, float64]{
		Left[float64, float64]{3.0},
		Left[float64, float64]{-3.9},
		Left[float64, float64]{3.1415},
	},
		[]float64{}

	in3, ex3 := []Either[float64, float64]{
		Right[float64, float64]{3.0},
		Right[float64, float64]{-3.9},
		Right[float64, float64]{3.1415},
	},
		[]float64{3, -3.9, 3.1415}

	test(ex1, Rights(in1), in1)
	test(ex2, Rights(in2), in2)
	test(ex3, Rights(in3), in3)
}

func TestPartitionEither(t *testing.T) {
	test := utils.AsrtCallEq(t, "PartitionEither")

	in1, ex1l, ex1r := []Either[int, string]{
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
		[]int{3, 4, 0, -912, 3099},
		[]string{
			"daniuh",
			"98u9",
			"",
			"synactic sugar causes cancer of the semicolon",
			"when will java die out already",
			"public static void main(String[] args)",
		}

	in2, ex2l, ex2r := []Either[float64, float64]{
		Left[float64, float64]{3.0},
		Left[float64, float64]{-3.9},
		Left[float64, float64]{3.1415},
	},
		[]float64{3, -3.9, 3.1415},
		[]float64{}

	in3, ex3l, ex3r := []Either[float64, float64]{
		Right[float64, float64]{3.0},
		Right[float64, float64]{-3.9},
		Right[float64, float64]{3.1415},
	},
		[]float64{},
		[]float64{3, -3.9, 3.1415}

	test(Tuple2[[]int, []string]{ex1l, ex1r}, PartitionEither(in1), in1)
	test(Tuple2[[]float64, []float64]{ex2l, ex2r}, PartitionEither(in2), in2)
	test(Tuple2[[]float64, []float64]{ex3l, ex3r}, PartitionEither(in3), in3)
}
