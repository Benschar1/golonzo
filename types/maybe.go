package types

import (
	"log"

	"github.com/Benschar1/golonzo/utils"
)

type Maybe[A any] interface {
	isMaybe()
}

type Some[A any] struct {
	val A
}

type None[A any] struct{}

func (_ Some[A]) isMaybe() {}
func (_ None[A]) isMaybe() {}

// IsSome :: Maybe a -> bool
func IsSome[A any](m Maybe[A]) bool {
	switch m.(type) {
	case Some[A]:
		return true
	default:
		return false
	}
}

// IsNone :: Maybe a -> bool
func IsNone[A any](m Maybe[A]) bool {
	switch m.(type) {
	case None[A]:
		return true
	default:
		return false
	}
}

// FromMaybe :: a -> Maybe a -> a
func FromMaybe[A any](df A) func(Maybe[A]) A {
	return func(m Maybe[A]) A {
		switch v := m.(type) {
		case None[A]:
			return df
		case Some[A]:
			return v.val
		default:
			log.Panic(utils.BadSumTypeConstructor(v, "Maybe", "Some", "None"))
			return Some[A]{}.val
		}
	}
}

// MapMaybe :: (a -> b) -> Maybe a -> Maybe b
func MapMaybe[A, B any](f func(A) B) func(Maybe[A]) Maybe[B] {
	return func(m Maybe[A]) Maybe[B] {
		switch v := m.(type) {
		case Some[A]:
			return Some[B]{f(v.val)}
		case None[A]:
			return None[B]{}
		default:
			log.Panic(utils.BadSumTypeConstructor(v, "Maybe", "Some", "None"))
			return None[B]{}
		}
	}
}

// BindMaybe :: Maybe a -> (a -> Maybe b) -> Maybe b
func BindMaybe[A, B any](m Maybe[A]) func(func(A) Maybe[B]) Maybe[B] {
	return func(f func(A) Maybe[B]) Maybe[B] {
		switch v := m.(type) {
		case Some[A]:
			return f(v.val)
		case None[A]:
			return None[B]{}
		default:
			log.Panic(utils.BadSumTypeConstructor(v, "Maybe", "Some", "None"))
			return None[A]{}
		}
	}
}

// FilterMaybes :: [Maybe a] -> [a]
func FilterMaybes[A any](ms []Maybe[A]) []A {
	vals := make([]A, 0, len(ms))
	for _, m := range ms {
		switch v := m.(type) {
		case Some[A]:
			vals = append(vals, v.val)
		case None[A]:
		default:
			log.Panic(utils.BadSumTypeConstructor(v, "Maybe", "Some", "None"))
		}
	}
	return vals
}
