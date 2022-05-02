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
