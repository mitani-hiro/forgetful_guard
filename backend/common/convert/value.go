package convert

// ToPointer プリミティブ型の値をポインターに変換する.
func ToPointer[T comparable](v T) *T {
	var zero T

	if v == zero {
		return nil
	}

	return &v
}
