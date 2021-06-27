package lib

func keys(m map[string]interface{}) []string {
	var ks []string
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}

func values(m map[string]interface{}) []interface{} {
	var vs []interface{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func multiKeys(ms ...map[string]interface{}) []string {
	var ks []string
	for _, m := range ms {
		ks = append(ks, keys(m)...)
	}
	return ks
}

func multiValues(ms ...map[string]interface{}) []interface{} {
	var vs []interface{}
	for _, m := range ms {
		vs = append(vs, values(m)...)
	}
	return vs
}

func get(m map[string]interface{}, k string) interface{} {
	if v, ok := m[k]; ok {
		return v
	}
	return ""
}

func set(m map[string]interface{}, k string, v interface{}) interface{} {
	m[k] = v
	return v
}

func del(m map[string]interface{}, k string) map[string]interface{} {
	delete(m, k)
	return m
}

func hasKey(m map[string]interface{}, k string) bool {
	_, ok := m[k]
	return ok
}

func dict(v ...interface{}) map[interface{}]interface{} {
	dict := map[interface{}]interface{}{}
	length := len(v)
	for i := 0; i < length; i += 2 {
		key := v[i]
		if i+1 >= length {
			dict[i] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}
