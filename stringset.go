package stringset

import (
	"encoding/json"
	"sync"
)

// New will create a new instance of StringSet
func New(keys ...string) *StringSet {
	var s StringSet
	s.m = makeMap(keys)
	return &s
}

// StringSet is a thread-safe wrapper for the stringset.Map
type StringSet struct {
	mux sync.RWMutex
	m   Map
}

// Set will insert a single key
func (s *StringSet) Set(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m.Set(key)
}

// SetMany will insert multiple keys
func (s *StringSet) SetMany(keys ...string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m.SetMany(keys)
}

// Unset will remove a key
func (s *StringSet) Unset(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m.Unset(key)
}

// UnsetMany will remove multiple keys
func (s *StringSet) UnsetMany(keys ...string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m.UnsetMany(keys)
}

// Has will return whether or not a key exists
func (s *StringSet) Has(key string) (has bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.m.Has(key)
}

// HasAll will return whether or not all the provided keys exist
func (s *StringSet) HasAll(keys ...string) (has bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.m.HasAll(keys)
}

// HasOne will return whether or not at least one of the provided keys exist
func (s *StringSet) HasOne(keys ...string) (has bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.m.HasOne(keys)
}

// Slice will return the keys as a slice
func (s *StringSet) Slice() (keys []string) {
	if s == nil {
		return
	}

	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.m.Slice()
}

// IsMatch will return whether or not two StringSets are an identical match
func (s *StringSet) IsMatch(a *StringSet) (isMatch bool) {
	if s == nil {
		return
	}

	s.mux.RLock()
	defer s.mux.RUnlock()
	a.mux.RLock()
	defer a.mux.RUnlock()

	return s.m.IsMatch(a.m)
}

// Len will return the length of the string set
func (s *StringSet) Len() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return len(s.m)
}

// MarshalJSON is a JSON encoding helper func
func (s *StringSet) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(s.m)
}

// UnmarshalJSON is a JSON decoding helper func
func (s *StringSet) UnmarshalJSON(bs []byte) (err error) {
	return json.Unmarshal(bs, &s.m)
}
