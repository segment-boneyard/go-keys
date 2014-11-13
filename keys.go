package keys

// Return keys in the given map.
func Keys(m map[string]interface{}) []string {
	var keys []string

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}
