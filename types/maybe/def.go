package main

type Maybe[A any] interface {
	isMaybe()
}

type Some[A any] struct{ val A }

type None[A any] struct{}

func (m Some[A]) isMaybe() {}
func (m None[A]) isMaybe() {}

func Match[A, B any](
	caseSome func(A) B,
	caseNone func() B,
) func(maybe Maybe[A]) B {
	return func(maybe Maybe[A]) B {
		switch t := maybe.(type) {
		case Some[A]:
			return caseSome(t.val)
		case None[A]:
			return caseNone()
		default:
			panic("expected Some or None but received neither")
		}
	}
}

func MatchDo[A any](
	caseSome func(A),
	caseNone func(),
) func(maybe Maybe[A]) {
	return func(maybe Maybe[A]) {
		switch t := maybe.(type) {
		case Some[A]:
			caseSome(t.val)
		case None[A]:
			caseNone()
		default:
			panic("expected Some or None but received neither")
		}
	}
}

func IsSome[A any](m Maybe[A]) bool {
	return Match(
		func(a A) bool { return true },
		func() bool { return false },
	)(m)
}

func IsNone[A any](m Maybe[A]) bool {
	return Match(
		func(a A) bool { return false },
		func() bool { return true },
	)(m)
}
