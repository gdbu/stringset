package stringset

import "testing"

func BenchmarkMap_Set(b *testing.B) {
	m := make(Map)
	for i := 0; i < b.N; i++ {
		m.Set("foo")
	}
}

func BenchmarkMap_Unset(b *testing.B) {
	m := make(Map)
	for i := 0; i < b.N; i++ {
		m.Unset("foo")
	}
}
