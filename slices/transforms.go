package slices

// Map :: (a -> b) -> []a -> []b
func Map[A, B any](f func(A) B) func([]A) []B {
	return func(sliceIn []A) []B {
		sliceOut := make([]B, len(sliceIn))

		for i, v := range sliceIn {
			sliceOut[i] = f(v)
		}

		return sliceOut
	}
}

// Reverse :: []a -> []a
func Reverse[A any](fwd []A) []A {
	bwd := make([]A, 0, len(fwd))

	for i := len(fwd) - 1; i >= 0; i-- {
		bwd = append(bwd, fwd[i])
	}

	return bwd
}

// Intersperse :: a -> []a -> []a
func Intersperse[A any](el A) func([]A) []A {
	return func(sliceIn []A) []A {
		if len(sliceIn) < 2 {
			return sliceIn
		}

		sliceOut := make([]A, 0, 2*len(sliceIn)-1)
		sliceOut = append(sliceOut, sliceIn[0])

		for i := 1; i < len(sliceIn); i++ {
			sliceOut = append(sliceOut, el, sliceIn[i])
		}

		return sliceOut
	}
}

// Intercalate :: []a -> [][]a -> []a
func Intercalate[A any](xs []A) func([][]A) []A {
	return func(xss [][]A) []A {
		return Concat(Intersperse(xs)(xss))
	}
}

// ConcatMap :: (a -> []b) -> []a -> []b
func ConcatMap[A, B any](f func(A) []B) func([]A) []B {
	return func(sliceIn []A) []B {
		return Concat(Map(f)(sliceIn))
	}
}
