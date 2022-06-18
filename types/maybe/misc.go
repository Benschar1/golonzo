package main

func ValOrDefault[A any](a A) func(ma Maybe[A]) A {
	return Match(
		func(aa A) A { return aa },
		func() A { return a },
	)
}

func FilterNones[A any](ms []Maybe[A]) []A {
	arr := make([]A, 0, len(ms))

	for _, ma := range ms {
		MatchDo(
			func(a A) { arr = append(arr, a) },
			func() {},
		)(ma)
	}

	return arr
}
