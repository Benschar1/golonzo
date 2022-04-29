package types

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Benschar1/golonzo/utils"
)

type Either[A, B any] interface {
	isEither()
}

type Left[A, B any] struct {
	Val A
}

type Right[A, B any] struct {
	Val B
}

func (Left[A, B]) isEither()  {}
func (Right[A, B]) isEither() {}

// IsLeft :: Either a b -> bool
func IsLeft[A, B any](e Either[A, B]) bool {
	switch e.(type) {
	case Left[A, B]:
		return true
	case Right[A, B]:
		return false
	default:
		return eitherSwitchDefault[bool](e)
	}
}

// IsRight :: Either a b -> bool
func IsRight[A, B any](e Either[A, B]) bool {
	switch e.(type) {
	case Right[A, B]:
		return true
	case Left[A, B]:
		return false
	default:
		return eitherSwitchDefault[bool](e)
	}
}

// FromLeft :: a -> Either a b -> a
func FromLeft[A, B any](df A) func(e Either[A, B]) A {
	return func(e Either[A, B]) A {
		switch v := e.(type) {
		case Left[A, B]:
			return v.Val
		case Right[A, B]:
			return df
		default:
			return eitherSwitchDefault[A](e)
		}
	}
}

// FromRight :: b -> Either a b -> b
func FromRight[A, B any](df B) func(e Either[A, B]) B {
	return func(e Either[A, B]) B {
		switch v := e.(type) {
		case Right[A, B]:
			return v.Val
		case Left[A, B]:
			return df
		default:
			return eitherSwitchDefault[B](e)
		}
	}
}

// EitherMapL :: (a1 -> a2) -> Either a1 b -> Either a2 b
func EitherMapL[A1, A2, B any](f func(A1) A2) func(e Either[A1, B]) Either[A2, B] {
	return func(e Either[A1, B]) Either[A2, B] {
		switch v := e.(type) {
		case Left[A1, B]:
			return Left[A2, B]{f(v.Val)}
		case Right[A1, B]:
			return Right[A2, B](v)
		default:
			return eitherSwitchDefault[Either[A2, B]](e)
		}
	}
}

// EitherMapR :: (b1 -> b2) -> Either a b1 -> Either a b2
func EitherMapR[B1, B2, A any](f func(B1) B2) func(e Either[A, B1]) Either[A, B2] {
	return func(e Either[A, B1]) Either[A, B2] {
		switch v := e.(type) {
		case Right[A, B1]:
			return Right[A, B2]{f(v.Val)}
		case Left[A, B1]:
			return Left[A, B2](v)
		default:
			return eitherSwitchDefault[Either[A, B2]](e)
		}
	}
}

// EitherMapLR :: (a1 -> a2) -> (b1 -> b2) -> Either a1 b1 -> Either a1 b2
func EitherMapLR[A1, A2, B1, B2 any](fl func(A1) A2) func(func(B1) B2) func(Either[A1, B1]) Either[A2, B2] {
	return func(fr func(B1) B2) func(Either[A1, B1]) Either[A2, B2] {
		return func(e Either[A1, B1]) Either[A2, B2] {
			switch v := e.(type) {
			case Right[A1, B1]:
				return Right[A2, B2]{fr(v.Val)}
			case Left[A1, B1]:
				return Left[A2, B2]{fl(v.Val)}
			default:
				return eitherSwitchDefault[Either[A2, B2]](e)
			}
		}
	}
}

// LeftToMaybe :: Either a b -> Maybe a
func LeftToMaybe[A, B any](e Either[A, B]) Maybe[A] {
	switch v := e.(type) {
	case Left[A, B]:
		return Some[A](v)
	case Right[A, B]:
		return None[A]{}
	default:
		return eitherSwitchDefault[Maybe[A]](e)
	}
}

// RightToMaybe :: Either a b -> Maybe b
func RightToMaybe[A, B any](e Either[A, B]) Maybe[B] {
	switch v := e.(type) {
	case Right[A, B]:
		return Some[B](v)
	case Left[A, B]:
		return None[B]{}
	default:
		return eitherSwitchDefault[Maybe[B]](e)
	}
}

// UnifyEither :: Either a a -> a
func UnifyEither[A any](e Either[A, A]) A {
	switch v := e.(type) {
	case Right[A, A]:
		return v.Val
	case Left[A, A]:
		return v.Val
	default:
		return eitherSwitchDefault[A](e)
	}
}

// EitherFlip :: Either a b -> Either b a
func EitherFlip[A, B any](e Either[A, B]) Either[B, A] {
	switch v := e.(type) {
	case Left[A, B]:
		return Right[B, A](v)
	case Right[A, B]:
		return Left[B, A](v)
	default:
		return eitherSwitchDefault[Either[B, A]](e)
	}
}

// Lefts :: [Either a b] -> [a]
func Lefts[A, B any](es []Either[A, B]) []A {
	lefts := make([]A, 0, len(es))
	for _, v := range es {
		switch v := v.(type) {
		case Left[A, B]:
			lefts = append(lefts, v.Val)
		case Right[A, B]:
		default:
			return eitherSwitchDefault[[]A](v)
		}
	}
	return lefts
}

// Rights :: [Either a b] -> [b]
func Rights[A, B any](es []Either[A, B]) []B {
	rights := make([]B, 0, len(es))
	for _, v := range es {
		switch v := v.(type) {
		case Right[A, B]:
			rights = append(rights, v.Val)
		case Left[A, B]:
		default:
			return eitherSwitchDefault[[]B](v)
		}
	}
	return rights
}

// PartitionEither :: [Either a b] -> ([a], [b])
func PartitionEither[A, B any](es []Either[A, B]) Tuple2[[]A, []B] {
	lefts := make([]A, 0, len(es))
	rights := make([]B, 0, len(es))
	for _, v := range es {
		switch v := v.(type) {
		case Left[A, B]:
			lefts = append(lefts, v.Val)
		case Right[A, B]:
			rights = append(rights, v.Val)
		default:
			return eitherSwitchDefault[Tuple2[[]A, []B]](v)
		}
	}
	return Tuple2[[]A, []B]{lefts, rights}
}

// error utilities

func EitherTypeSig[A, B any]() string {
	return fmt.Sprintf(
		"Either[%[1]s, %[2]s] :: Left{ Val %[1]s } | Right{ Val %[2]s }",
		reflect.TypeOf((*A)(nil)).Elem(),
		reflect.TypeOf((*B)(nil)).Elem(),
	)
}

func EitherTypeError[A, B any](e Either[A, B]) string {
	return utils.BadTypeError(e, EitherTypeSig[A, B]())
}

func eitherSwitchDefault[Ret, A, B any](e Either[A, B]) Ret {
	log.Panicf(EitherTypeError(e))
	var x Ret
	return x
}
