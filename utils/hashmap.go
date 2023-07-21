package utils

func MapKeys[K comparable, V any](data map[K]V) []K {
	result := make([]K, 0, len(data))
	for key := range data {
		result = append(result, key)
	}
	return result
}

func MapValues[K comparable, V any](data map[K]V) []V {
	result := make([]V, 0, len(data))
	for _, value := range data {
		result = append(result, value)
	}
	return result
}
