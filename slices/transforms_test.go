package slices

import (
	"testing"

	"github.com/Benschar1/golonzo/utils"
)

func TestMap(t *testing.T) {
	test := utils.AsrtCallEq(t, "Map")

	f1, in1, ex1 :=
		func(str string) string { return str + "abc" },
		[]string{"af", "", "what does the fox say", "john coltrane was tone deaf"},
		[]string{"afabc", "abc", "what does the fox sayabc", "john coltrane was tone deafabc"}

	f2, in2, ex2 :=
		func(num int) bool { return num > 0 },
		[]int{1, 3, 5, -2, -20, 0},
		[]bool{true, true, true, false, false, false}

	test(ex1, Map(f1)(in1), f1, in1)
	test(ex2, Map(f2)(in2), f2, in2)
}

func TestReverse(t *testing.T) {
	test := utils.AsrtCallEq(t, "Reverse")

	in1, ex1 := []int{1, 3, 7, -2, 209}, []int{209, -2, 7, 3, 1}
	in2, ex2 := []float64{387, -3.5}, []float64{-3.5, 387}
	in3, ex3 := []bool{true, false, false}, []bool{false, false, true}
	in4, ex4 := []int{}, []int{}
	in5, ex5 := []int{1}, []int{1}

	test(ex1, Reverse(in1), in1)
	test(ex2, Reverse(in2), in2)
	test(ex3, Reverse(in3), in3)
	test(ex4, Reverse(in4), in4)
	test(ex5, Reverse(in5), in5)
}

func TestIntersperse(t *testing.T) {
	test := utils.AsrtCallEq(t, "Intersperse")

	el1, arr1, ex1 := 3, []int{1, 7, 9, -5}, []int{1, 3, 7, 3, 9, 3, -5}
	el2, arr2, ex2 := "xhj", []string{"g"}, []string{"g"}
	el3, arr3, ex3 := 4, []int{}, []int{}

	test(ex1, Intersperse(el1)(arr1), el1, arr1)
	test(ex2, Intersperse(el2)(arr2), el2, arr2)
	test(ex3, Intersperse(el3)(arr3), el3, arr3)
}

func TestIntercalate(t *testing.T) {
	test := utils.AsrtCallEq(t, "Intercalate")

	xs1, xss1, ex1 :=
		[]int{2, 5},
		[][]int{{1, 2}, {3, 4}, {5, 6}},
		[]int{1, 2, 2, 5, 3, 4, 2, 5, 5, 6}

	xs2, xss2, ex2 :=
		[]rune("xhj"),
		[][]rune{[]rune("g")},
		[]rune("g")

	xs3, xss3, ex3 :=
		[]int{4},
		[][]int{},
		[]int{}

	test(ex1, Intercalate(xs1)(xss1), xs1, xss1)
	test(ex2, Intercalate(xs2)(xss2), xs2, xss2)
	test(ex3, Intercalate(xs3)(xss3), xs3, xss3)
}

func TestConcatMap(t *testing.T) {
	test := utils.AsrtCallEq(t, "ConcatMap")

	f1, s1, ex1 :=
		Replicate[int](3),
		[]int{1, 3, 2, 10},
		[]int{1, 1, 1, 3, 3, 3, 2, 2, 2, 10, 10, 10}

	f2, s2, ex2 :=
		func(s string) []int { return []int{} },
		[]string{"dj", "hjh", "uiy"},
		[]int{}

	f3, s3, ex3 :=
		func(i int) []int { return Replicate[int](i)(i) },
		[]int{-3, 0, -2, 2, -9, 4, 0, 1},
		[]int{2, 2, 4, 4, 4, 4, 1}

	test(ex1, ConcatMap(f1)(s1), f1, s1)
	test(ex2, ConcatMap(f2)(s2), f2, s2)
	test(ex3, ConcatMap(f3)(s3), f3, s3)
}
