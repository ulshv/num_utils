package num_utils

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

func Add[T Number](a, b T) T {
	return a + b
}

func Mul[T Number](a, b T) T {
	return a * b
}
