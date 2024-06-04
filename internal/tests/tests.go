package tests

import "strings"

func StringToMap(s string) map[string]struct{} {
	m := make(map[string]struct{})
	parts := strings.Split(s, " ")
	for i := 0; i < len(parts); i++ {
		m[parts[i]] = struct{}{}
	}
	return m
}

func MapsEqual(m1, m2 map[string]struct{}) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			println("key not found: ", k)
			return false
		}
	}
	return true
}
