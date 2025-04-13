package utils

func ToPointer[T any](t T) *T {
	return &t
}

func FromPointer[T any](t *T) T {
	return *t
}
