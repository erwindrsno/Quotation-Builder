package util

func MapList[T any, V any](items []T, mapper func(T) V) []V {
	result := make([]V, len(items))

	for i, item := range items {
		result[i] = mapper(item)
	}

	return result
}
