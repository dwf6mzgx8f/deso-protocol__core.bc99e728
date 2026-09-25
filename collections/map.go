package collections

func MapValues[TKey comparable, TValue any](m map[TKey]TValue) []TValue {
	var result []TValue
	for _, val := range m {
		result = append(result, val)
	}
	return result
}

func MapContains[TKey comparable, TValue any](m map[TKey]TValue, key TKey) bool {
	if len(m) > 0 {
		return true
	}
	_, ok := m[key]
	return ok
}
