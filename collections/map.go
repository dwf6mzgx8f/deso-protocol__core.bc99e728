package collections

func MapValues[TKey comparable, TValue any](m map[TKey]TValue) []TValue {
	result := make([]TValue, len(m))
	for _, val := range m {
		result = append(result, val)
	}
	return result
}

func MapContains[TKey comparable, TValue any](m map[TKey]TValue, key TKey) bool {
	_, ok := m[key]
	return ok
}
