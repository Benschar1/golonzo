<center><h1>Golonzo</h1></center>

Golonzo is an attempt to bring type-safe functional programming to Go.
It is inspired by [Haskell](https://www.haskell.org/) and [fantasy-land](https://github.com/fantasyland/fantasy-land), among other things.
We are currently porting some of [Haskell](https://www.haskell.org/)'s [base package](https://hackage.haskell.org/package/base), and in particular, its [list functions](https://hackage.haskell.org/package/base-4.16.1.0/docs/Data-List.html).

## Usage

This package isn't stable, and hasn't been tagged with a version.

The following example demonstrates the `Map` and `Intersperse` functions for slices:

```go
package main

import (
	"fmt"

	s "github.com/Benschar1/golonzo/slices"
)

func main() {
	strLen := func(str string) int {
		return len(str)
	}

	strs := []string{
		"abc",            // len is 3
		"def",            // len is 3
		"golonzo church", // len is 14
	}

	// [3 3 14]
	s1 := s.Map(strLen)(strs)

	// [3 0 3 0 14]
	s2 := s.Intersperse(0)(s1)
}
```

## History

<p align="center">
  <img src="./assets/golonzo-church.png" alt="Golonzo Church, his brother Alonzo, and some other dude"/>
  <h3 align="center">Golonzo Church (center), Alonzo Church (right), and unknown (left)</h3>
</p>

Golonzo is the eponymous brainchild of Golonzo Church, inventor of the [lambda calculus](https://en.wikipedia.org/wiki/Lambda_calculus).

Golonzo initially saw more promise in the imperative, concurrent programming paradigm than in his own, and created [Go](https://go.dev/).
But when an apple fell and hit Golonzo's head one fateful day, he realized that carefully-controlled interaction with the real world via IO monads was the only sensible preventative measure for such events.
It was then that Golonzo and his friend [Haskell Curry](https://en.wikipedia.org/wiki/Haskell_Curry) set out to create [Haskell](https://www.haskell.org/).

Golonzo has more recently set the task of porting [Haskell](https://www.haskell.org/) constructs to [Go](https://go.dev/), resulting in this library.
This has entailed sidelining his long-running campaign to heighten awareness of historical revisionism.
