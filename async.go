package stringset

import (
	"encoding/json"
	"sync"
)

var _ StringSet = &Async{}

// NewAsyncMap will create a new Map instance of Async
func NewAsyncMap(keys ...string) *Async {
	var a Async
	a.s = newMap(keys)
	return &a
}

// Async is a thread-safe wrapper for the stringset.Map
type Async struct {
	mux sync.RWMutex
	s   StringSet
}

// Set will insert a single key
func (a *Async) Set(key string) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.s.Set(key)
}

// SetMany will insert multiple keys
func (a *Async) SetMany(keys []string) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.s.SetMany(keys)
}

// SetMany will insert multiple keys
func (a *Async) SetAsString(value string) (err error) {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.s.SetAsString(value)
}

// Unset will remove a key
func (a *Async) Unset(key string) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.s.Unset(key)
}

// UnsetMany will remove multiple keys
func (a *Async) UnsetMany(keys []string) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.s.UnsetMany(keys)
}

// Has will return whether or not a key exists
func (a *Async) Has(key string) (has bool) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.s.Has(key)
}

// HasAll will return whether or not all the provided keys exist
func (a *Async) HasAll(keys []string) (has bool) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.s.HasAll(keys)
}

// HasOne will return whether or not at least one of the provided keys exist
func (a *Async) HasOne(keys []string) (has bool) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.s.HasOne(keys)
}

// Slice will return the keys as a slice
func (a *Async) Slice() (keys []string) {
	if a == nil {
		return
	}

	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.s.Slice()
}

// IsMatch will return whether or not two Asyncs are an identical match
func (a *Async) IsMatch(in StringSet) (isMatch bool) {
	if a == nil {
		return
	}

	a.mux.RLock()
	defer a.mux.RUnlock()
	if inAsync, ok := in.(*Async); ok {
		inAsync.mux.RLock()
		defer inAsync.mux.RUnlock()
	}

	return a.s.IsMatch(in)
}

// Len will return the length of the string set
func (a *Async) Len() int {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.s.Len()
}

// MarshalJSON is a JSON encoding helper func
func (a *Async) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(a.s)
}

// UnmarshalJSON is a JSON decoding helper func
func (a *Async) UnmarshalJSON(bs []byte) (err error) {
	return json.Unmarshal(bs, &a.s)
}
