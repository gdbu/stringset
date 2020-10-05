package stringset

import "encoding/json"

func makeMap(keys []string) (m Map) {
	m = make(Map, len(keys))
	for _, key := range keys {
		m.Set(key)
	}

	return
}

// Map is the lower-level underlying type for storing a string set
// Note: This is not thread-safe. If you need thread-safety, please use StringSet type
type Map map[string]struct{}

// Set will place a key
func (m *Map) Set(key string) {
	mm := *m
	if mm == nil {
		*m = make(Map, 1)
		mm = *m
	}

	mm[key] = struct{}{}
}

// SetMany will place multiple keys
func (m Map) SetMany(keys []string) {
	for _, key := range keys {
		m.Set(key)
	}
}

// Unset will remove a key
func (m Map) Unset(key string) {
	if m == nil {
		return
	}

	delete(m, key)
}

// UnsetMany will remove multiple keys
func (m Map) UnsetMany(keys []string) {
	for _, key := range keys {
		m.Unset(key)
	}
}

// Has will return whether or not a key exists
func (m Map) Has(key string) (has bool) {
	if m == nil {
		return
	}

	_, has = m[key]
	return
}

// HasAll will return whether or not all the provided keys exist
func (m Map) HasAll(keys []string) (has bool) {
	for _, key := range keys {
		if !m.Has(key) {
			return
		}
	}

	return true
}

// HasOne will return whether or not at least one of the provided keys exist
func (m Map) HasOne(keys []string) (has bool) {
	for _, key := range keys {
		if m.Has(key) {
			return true
		}
	}

	return false
}

// Slice will return the keys as a slice
func (m Map) Slice() (keys []string) {
	if m == nil {
		return
	}

	for key := range m {
		keys = append(keys, key)
	}

	return
}

// MarshalJSON is a JSON encoding helper func
func (m Map) MarshalJSON() (bs []byte, err error) {
	return json.Marshal(m.Slice())
}

// UnmarshalJSON is a JSON decoding helper func
func (m *Map) UnmarshalJSON(bs []byte) (err error) {
	var keys []string
	if err = json.Unmarshal(bs, &keys); err != nil {
		return
	}

	*m = makeMap(keys)
	return
}
