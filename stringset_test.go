package stringset

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testKeys *StringSet

func TestStringSet_MarshalJSON(t *testing.T) {
	var (
		output []byte
		err    error
	)

	ss := New("foo", "bar", "baz")

	if output, err = ss.MarshalJSON(); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output!", string(output))
}

func TestStringSet_UnmarshalJSON(t *testing.T) {
	var (
		ss  *StringSet
		err error
	)

	if err = json.Unmarshal([]byte(`["foo","bar","baz"]`), &ss); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Parsed!", ss)
}

func ExampleNew() {
	// Initialize new stringset
	testKeys = New()
}

func ExampleStringSet_Set() {
	// Set foo key
	testKeys.Set("foo")
	// Set bar key
	testKeys.Set("bar")
}

func ExampleStringSet_Unset() {
	// Remove bar key
	testKeys.Unset("bar")
}

func ExampleStringSet_Has() {
	if testKeys.Has("foo") {
		fmt.Println("We have foo!")
	}

	if !testKeys.Has("bar") {
		fmt.Println("We do not have bar!")
	}
}

func ExampleStringSet_Slice() {
	keys := testKeys.Slice()
	for _, key := range keys {
		fmt.Printf("Iterating key: %s\n", key)
	}
}
