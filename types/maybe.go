package types

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Benschar1/golonzo/utils"
)

type Maybe[A any] interface {
	isMaybe()
}

type Some[A any] struct {
	Val A
}

type None[A any] struct{}

func (Some[A]) isMaybe() {}
func (None[A]) isMaybe() {}

// IsSome :: Maybe a -> bool
func IsSome[A any](m Maybe[A]) bool {
	switch m.(type) {
	case Some[A]:
		return true
	case None[A]:
		return false
	default:
		return maybeSwitchDefault[bool](m)
	}
}

// IsSome :: Maybe a -> bool
func IsNone[A any](m Maybe[A]) bool {
	switch m.(type) {
	case None[A]:
		return true
	case Some[A]:
		return false
	default:
		return maybeSwitchDefault[bool](m)
	}
}

// FromMaybe :: a -> Maybe a -> a
func FromMaybe[A any](df A) func(Maybe[A]) A {
	return func(m Maybe[A]) A {
		switch v := m.(type) {
		case None[A]:
			return df
		case Some[A]:
			return v.Val
		default:
			return maybeSwitchDefault[A](m)
		}
	}
}

// MapMaybe :: (a -> b) -> Maybe a -> Maybe b
func MapMaybe[A, B any](f func(A) B) func(Maybe[A]) Maybe[B] {
	return func(m Maybe[A]) Maybe[B] {
		switch v := m.(type) {
		case Some[A]:
			return Some[B]{f(v.Val)}
		case None[A]:
			return None[B]{}
		default:
			return maybeSwitchDefault[Maybe[B]](m)
		}
	}
}

// BindMaybe :: Maybe a -> (a -> Maybe b) -> Maybe b
func BindMaybe[A, B any](m Maybe[A]) func(func(A) Maybe[B]) Maybe[B] {
	return func(f func(A) Maybe[B]) Maybe[B] {
		switch v := m.(type) {
		case Some[A]:
			return f(v.Val)
		case None[A]:
			return None[B]{}
		default:
			return maybeSwitchDefault[Maybe[B]](m)
		}
	}
}

// FilterMaybes :: []Maybe a -> []a
func FilterMaybes[A any](ms []Maybe[A]) []A {
	vals := make([]A, 0, len(ms))
	for _, m := range ms {
		switch v := m.(type) {
		case Some[A]:
			vals = append(vals, v.Val)
		case None[A]:
		default:
			return maybeSwitchDefault[[]A](m)
		}
	}
	return vals
}

// error utilities

func MaybeTypeSig[A any]() string {
	return fmt.Sprintf(
		"Maybe[%s] :: Some{ Val %s } | None",
		reflect.TypeOf((*A)(nil)).Elem(),
		reflect.TypeOf((*A)(nil)).Elem(),
	)
}

func MaybeTypeError[A any](m Maybe[A]) string {
	return utils.BadTypeError(m, MaybeTypeSig[A]())
}

func maybeSwitchDefault[Ret, A any](m Maybe[A]) Ret {
	log.Panicf(MaybeTypeError(m))
	var x Ret
	return x
}
