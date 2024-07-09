package stringset

import (
	"encoding/json"
	"strings"

	"github.com/gdbu/bst"
)

var _ StringSet = &Map{}

// MakeMap will initialize a new map
func NewSorted(keys ...string) *Sorted {
	s := makeSorted(keys)
	return &s
}

func makeSorted(keys []string) (s Sorted) {
	s.d = *bst.NewKeys(keys...)
	return
}

// Sorted is a sorted Stringset
type Sorted struct {
	d bst.Keys
}

// Set will place a key
func (s *Sorted) Set(key string) {
	s.d.Set(key)
}

// SetMany will place multiple keys
func (s *Sorted) SetMany(keys []string) {
	for _, key := range keys {
		s.Set(key)
	}
}

// Unset will remove a key
func (s *Sorted) Unset(key string) {
	if s == nil {
		return
	}

	s.d.Unset(key)
}

// UnsetMany will remove multiple keys
func (s *Sorted) UnsetMany(keys []string) {
	for _, key := range keys {
		s.Unset(key)
	}
}

// Has will return whether or not a key exists
func (s *Sorted) Has(key string) (has bool) {
	if s == nil {
		return
	}

	return s.d.Has(key)
}

// HasAll will return whether or not all the provided keys exist
func (s *Sorted) HasAll(keys []string) (has bool) {
	for _, key := range keys {
		if !s.Has(key) {
			return
		}
	}

	return true
}

// HasOne will return whether or not at least one of the provided keys exist
func (s *Sorted) HasOne(keys []string) (has bool) {
	for _, key := range keys {
		if s.Has(key) {
			return true
		}
	}

	return false
}

// Slice will return the keys as a slice
func (s *Sorted) Slice() (keys []string) {
	if s == nil {
		return
	}

	return s.d.Slice()
}

// IsMatch will return whether or not two Maps are an identical match
func (s *Sorted) IsMatch(in StringSet) (isMatch bool) {
	if s.d.Len() != in.Len() {
		return false
	}

	isMatch = !s.d.ForEach(func(key string) (end bool) {
		return !in.Has(key)
	})

	return
}

// Len will return the length of the map
func (s *Sorted) Len() (n int) {
	if s == nil {
		return
	}

	return s.d.Len()
}

// MarshalJSON is a JSON encoding helper func
func (s *Sorted) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(s.Slice())
}

// UnmarshalJSON is a JSON decoding helper func
func (s *Sorted) UnmarshalJSON(bs []byte) (err error) {
	var keys []string
	if err = json.Unmarshal(bs, &keys); err != nil {
		return
	}

	*s = makeSorted(keys)
	return
}

// SetAsString will set a map from a comma separated string value
// Note: This allows stringset to match the reflectio.Setter interface
func (s *Sorted) SetAsString(value string) (err error) {
	spl := strings.Split(value, ",")
	s.SetMany(spl)
	return
}
