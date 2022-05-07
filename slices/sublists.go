package slices

// Take :: int -> []a -> []a
func Take[A any](i int) func([]A) []A {
	return func(slice []A) []A {
		switch {
		case i < 1:
			return []A{}
		case i >= len(slice):
			return slice
		default:
			return slice[:i]
		}
	}
}

// Drop :: int -> []a -> []a
func Drop[A any](i int) func([]A) []A {
	return func(slice []A) []A {
		switch {
		case i < 1:
			return slice
		case i >= len(slice):
			return []A{}
		default:
			return slice[i:]
		}
	}
}

// Filter :: (a -> bool) -> []a -> []a
func Filter[A any](pred func(A) bool) func([]A) []A {
	return func(slice []A) []A {
		out := make([]A, 0, len(slice))
		for _, v := range slice {
			if pred(v) {
				out = append(out, v)
			}
		}

		return out
	}
}
