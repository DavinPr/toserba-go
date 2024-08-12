package common

func ToPtr[T any](val T) *T {
	return &val
}

func FromPtr[T any](val *T) (t T) {
	if val != (*T)(nil) {
		ptr := *val
		return ptr
	}

	return t
}
