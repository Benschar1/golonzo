package slices

import (
	"testing"

	"github.com/Benschar1/golonzo/utils"
)

func TestTake(t *testing.T) {

	test := utils.AsrtCallEq(t, "Take")

	n1, s1, ex1 :=
		3,
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3}

	n2, s2, ex2 :=
		1,
		[]string{"a", "b", "c"},
		[]string{"a"}

	n3, s3, ex3 :=
		0,
		[]int{4, 5, 1, 0, -10},
		[]int{}

	n4, s4, ex4 :=
		-1,
		[]int{4, 5, 1, 0, -10},
		[]int{}

	n5, s5, ex5 :=
		6,
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c"}

	test(ex1, Take[int](n1)(s1), n1, s1)
	test(ex2, Take[string](n2)(s2), n2, s2)
	test(ex3, Take[int](n3)(s3), n3, s3)
	test(ex4, Take[int](n4)(s4), n4, s4)
	test(ex5, Take[string](n5)(s5), n5, s5)
}

func TestDrop(t *testing.T) {

	test := utils.AsrtCallEq(t, "Drop")

	n1, s1, ex1 :=
		3,
		[]int{1, 2, 3, 4, 5},
		[]int{4, 5}

	n2, s2, ex2 :=
		1,
		[]string{"a", "b", "c"},
		[]string{"b", "c"}

	n3, s3, ex3 :=
		0,
		[]int{4, 5, 1, 0, -10},
		[]int{4, 5, 1, 0, -10}

	n4, s4, ex4 :=
		-1,
		[]int{4, 5, 1, 0, -10},
		[]int{4, 5, 1, 0, -10}

	n5, s5, ex5 :=
		6,
		[]string{"a", "b", "c"},
		[]string{}

	test(ex1, Drop[int](n1)(s1), n1, s1)
	test(ex2, Drop[string](n2)(s2), n2, s2)
	test(ex3, Drop[int](n3)(s3), n3, s3)
	test(ex4, Drop[int](n4)(s4), n4, s4)
	test(ex5, Drop[string](n5)(s5), n5, s5)
}

func TestFilter(t *testing.T) {

	test := utils.AsrtCallEq(t, "Filter")

	p1, s1, ex1 :=
		func(i int) bool { return i > 0 },
		[]int{1, -2, 3, -4, 5, -6, 7},
		[]int{1, 3, 5, 7}

	p2, s2, ex2 :=
		func(i int) bool { return true },
		[]int{1, -2, 3, -4, 5, -6, 7},
		[]int{1, -2, 3, -4, 5, -6, 7}

	p3, s3, ex3 :=
		func(i int) bool { return false },
		[]int{1, -2, 3, -4, 5, -6, 7},
		[]int{}

	test(ex1, Filter(p1)(s1), p1, s1)
	test(ex2, Filter(p2)(s2), p2, s2)
	test(ex3, Filter(p3)(s3), p3, s3)
}
