package collect

func Keys(m map[string]interface{}) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MultiKeys(ms ...map[string]interface{}) []string {
	var keys []string
	for _, m := range ms {
		keys = append(keys, Keys(m)...)
	}
	return keys
}
