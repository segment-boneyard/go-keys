package keys

import "sort"

// Keys of the given map.
func Keys(m map[string]interface{}) []string {
	var keys []string

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

// Sorted keys of the given map.
func Sorted(m map[string]interface{}) []string {
	keys := Keys(m)
	sort.Strings(keys)
	return keys
}
