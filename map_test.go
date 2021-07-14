package stringset

import (
	"fmt"
	"testing"
)

var (
	testBytes []byte
)

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

func BenchmarkMap_MarshalJSON(b *testing.B) {
	m := createTestSet(128)
	b.ResetTimer()

	var err error
	for i := 0; i < b.N; i++ {
		if testBytes, err = m.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkMap_UnmarshalJSON(b *testing.B) {
	var err error
	source := createTestSet(128)
	if testBytes, err = source.MarshalJSON(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var m Map
		if err = m.UnmarshalJSON(testBytes); err != nil {
			return
		}
	}

	b.ReportAllocs()
}

func createTestSet(size int) (m Map) {
	m = make(Map, size)
	for i := 0; i < size; i++ {

		key := fmt.Sprintf("%012d", i)
		m.Set(key)
	}

	return
}
