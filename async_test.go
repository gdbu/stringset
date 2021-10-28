package stringset

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testKeys *Async

func TestAsync_MarshalJSON(t *testing.T) {
	var (
		output []byte
		err    error
	)

	ss := NewAsyncMap("foo", "bar", "baz")

	if output, err = ss.MarshalJSON(); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output!", string(output))
}

func TestAsync_UnmarshalJSON(t *testing.T) {
	var (
		ss  *Async
		err error
	)

	if err = json.Unmarshal([]byte(`["foo","bar","baz"]`), &ss); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Parsed!", ss)
}

func ExampleNewAsyncMap() {
	// Initialize new stringset
	testKeys = NewAsyncMap()
}

func ExampleAsync_Set() {
	// Set foo key
	testKeys.Set("foo")
	// Set bar key
	testKeys.Set("bar")
}

func ExampleAsync_Unset() {
	// Remove bar key
	testKeys.Unset("bar")
}

func ExampleAsync_Has() {
	if testKeys.Has("foo") {
		fmt.Println("We have foo!")
	}

	if !testKeys.Has("bar") {
		fmt.Println("We do not have bar!")
	}
}

func ExampleAsync_Slice() {
	keys := testKeys.Slice()
	for _, key := range keys {
		fmt.Printf("Iterating key: %s\n", key)
	}
}
