package types

type Tuple0 struct{}

type Tuple1[T1 any] struct {
	V1 T1
}

type Tuple2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

type Tuple3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

type Tuple4[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

type Tuple5[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
}

type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
}

type Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
	V9 T9
}

type Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	V1  T1
	V2  T2
	V3  T3
	V4  T4
	V5  T5
	V6  T6
	V7  T7
	V8  T8
	V9  T9
	V10 T10
}
