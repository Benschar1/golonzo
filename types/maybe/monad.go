package main

// implement functor

func Map[A, B any](f func(A) B) func(ma Maybe[A]) Maybe[B] {
	return Match(
		func(a A) Maybe[B] { return Some[B]{f(a)} },
		func() Maybe[B] { return None[B]{} },
	)
}

func Fill[A, B any](ma Maybe[A]) func(b B) Maybe[B] {
	return func(b B) Maybe[B] {
		return Map(func(a A) B { return b })(ma)
	}
}

// implement applicative

func Pure[A any](a A) Maybe[A] {
	return Some[A]{a}
}

func Map2[A, B, C any](f func(A, B) C) func(ma Maybe[A]) func(mb Maybe[B]) Maybe[C] {
	return func(ma Maybe[A]) func(mb Maybe[B]) Maybe[C] {
		return func(mb Maybe[B]) Maybe[C] {
			return Match(
				func(a A) Maybe[C] {
					return Match(
						func(b B) Maybe[C] { return Some[C]{f(a, b)} },
						func() Maybe[C] { return None[C]{} },
					)(mb)
				},
				func() Maybe[C] { return None[C]{} },
			)(ma)
		}
	}
}

func Splat[A, B any](f Maybe[func(A) B]) func(m Maybe[A]) Maybe[B] {
	return Map2(
		func(f func(A) B, a A) B {
			return f(a)
		},
	)(f)
}

// implement monad

func Bind[A, B any](m Maybe[A]) func(f func(a A) Maybe[B]) Maybe[B] {
	return func(f func(a A) Maybe[B]) Maybe[B] {
		return Match(
			func(a A) Maybe[B] { return f(a) },
			func() Maybe[B] { return None[B]{} },
		)(m)
	}
}

func BindComp[A, B, C any](f1 func(A) Maybe[B]) func(f2 func(B) Maybe[C]) func(A) Maybe[C] {
	return func(f2 func(B) Maybe[C]) func(A) Maybe[C] {
		return func(a A) Maybe[C] {
			return Bind[B, C](f1(a))(f2)
		}
	}
}

func Join[A any](m Maybe[Maybe[A]]) Maybe[A] {
	return Match(
		func(a Some[A]) Maybe[A] { return a },
		func() Maybe[A] { return None[A]{} },
	)(m)
}
