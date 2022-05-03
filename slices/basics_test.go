package slices

import (
	"testing"

	tp "github.com/Benschar1/golonzo/types"
	"github.com/Benschar1/golonzo/utils"
)

func TestGet(t *testing.T) {
	test := utils.AsrtCallEq(t, "Get")

	arr1, ind1, ex1 := []string{"abc", "adf", "", "kj98d"}, 2, tp.Some[string]{""}

	arr2, ind2, ex2 := []int{387}, 2, tp.None[int]{}

	arr3, ind3, ex3 := []string{"abc", "adf", "", "kj98d"}, 4, tp.None[string]{}

	arr4, ind4, ex4 := []string{"abc", "adf", "", "kj98d"}, -34, tp.None[string]{}

	test(ex1, Get(arr1)(ind1), arr1, ind1)
	test(ex2, Get(arr2)(ind2), arr2, ind2)
	test(ex3, Get(arr3)(ind3), arr3, ind3)
	test(ex4, Get(arr4)(ind4), arr4, ind4)
}

func TestHead(t *testing.T) {
	test := utils.AsrtCallEq(t, "Head")

	arr1, ex1 := []string{"abc", "adf", "", "kj98d"}, tp.Some[string]{"abc"}
	arr2, ex2 := []string{}, tp.None[string]{}

	test(ex1, Head(arr1), arr1)
	test(ex2, Head(arr2), arr2)
}

func TestTail(t *testing.T) {
	test := utils.AsrtCallEq(t, "Tail")

	arr1, ex1 := []string{"abc", "adf", "", "kj98d"}, []string{"adf", "", "kj98d"}
	arr2, ex2 := []string{"f"}, []string{}
	arr3, ex3 := []string{}, []string{}

	test(ex1, Tail(arr1), arr1)
	test(ex2, Tail(arr2), arr2)
	test(ex3, Tail(arr3), arr3)
}

func TestInit(t *testing.T) {
	test := utils.AsrtCallEq(t, "Init")

	arr1, ex1 := []string{"abc", "adf", "", "kj98d"}, []string{"abc", "adf", ""}
	arr2, ex2 := []string{"f"}, []string{}
	arr3, ex3 := []string{}, []string{}

	test(ex1, Init(arr1), arr1)
	test(ex2, Init(arr2), arr2)
	test(ex3, Init(arr3), arr3)
}

func TestLast(t *testing.T) {
	test := utils.AsrtCallEq(t, "Head")

	arr1, ex1 := []string{"abc", "adf", "", "kj98d"}, tp.Some[string]{"kj98d"}
	arr2, ex2 := []string{}, tp.None[string]{}

	test(ex1, Last(arr1), arr1)
	test(ex2, Last(arr2), arr2)
}

func TestReplicate(t *testing.T) {
	test := utils.AsrtCallEq(t, "Replicate")

	num1, el1, ex1 := 4, "abc", []string{"abc", "abc", "abc", "abc"}
	num2, el2, ex2 := -2, 8, []int{}
	num3, el3, ex3 := 1, []int{3, 4}, [][]int{{3, 4}}

	test(ex1, Replicate[string](num1)(el1), el1, num1)
	test(ex2, Replicate[int](num2)(el2), el2, num2)
	test(ex3, Replicate[[]int](num3)(el3), el3, num3)

}

func TestConcat(t *testing.T) {
	test := utils.AsrtCallEq(t, "Concat")

	arr1, ex1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	arr2, ex2 := [][]string{}, []string{}

	arr3, ex3 := [][]string{
		{},
		{},
	},
		[]string{}

	test(ex1, Concat(arr1), arr1)
	test(ex2, Concat(arr2), arr2)
	test(ex3, Concat(arr3), arr3)
}

func TestEmpty(t *testing.T) {
	test := utils.AsrtCallEq(t, "Empty")

	arr1, ex1 := []int{1, 2}, false
	arr2, ex2 := []string{}, true
	arr3, ex3 := make([]int, 0, 10), true
	arr4, ex4 := make([]int, 10), false

	test(ex1, Empty(arr1), arr1)
	test(ex2, Empty(arr2), arr2)
	test(ex3, Empty(arr3), arr3)
	test(ex4, Empty(arr4), arr4)
}
