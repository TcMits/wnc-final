package generic

func Zero[T any]() T {
	var zero T
	return zero
}
func GetPointer[T any](v T) *T {
	return &v
}
