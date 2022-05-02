package slices

import t "github.com/Benschar1/golonzo/types"

// Get :: []a -> int -> Maybe a
func Get[A any](slice []A) func(int) t.Maybe[A] {
	return func(index int) t.Maybe[A] {
		if index < 0 || len(slice) <= index {
			return t.None[A]{}
		}
		return t.Some[A]{slice[index]}
	}
}

// Head :: []a -> Maybe a
func Head[A any](slice []A) t.Maybe[A] {
	return Get(slice)(0)
}

// Tail :: []a -> []a
func Tail[A any](slice []A) []A {
	if len(slice) <= 1 {
		return []A{}
	}
	return slice[1:]
}

// Init :: []a -> []a
func Init[A any](slice []A) []A {
	if len(slice) <= 1 {
		return []A{}
	}
	return slice[:len(slice)-1]
}

// Last :: []a -> Maybe a
func Last[A any](slice []A) t.Maybe[A] {
	return Get(slice)(len(slice) - 1)
}

// Replicate :: int -> a -> []a
func Replicate[A any](num int) func(A) []A {
	return func(el A) []A {
		if num < 1 {
			return []A{}
		}
		slice := make([]A, num)
		for i := range slice {
			slice[i] = el
		}
		return slice
	}
}

// Concat :: [][]a -> []a
func Concat[A any](slicesIn [][]A) []A {
	slicesOut := make([]A, 0)

	for _, v := range slicesIn {
		slicesOut = append(slicesOut, v...)
	}

	return slicesOut
}
