package stringset

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
func (m Map) Set(key string) {
	m[key] = struct{}{}
}

// SetMany will place multiple keys
func (m Map) SetMany(keys []string) {
	for _, key := range keys {
		m.Set(key)
	}
}

// Unset will remove a key
func (m Map) Unset(key string) {
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
	for key := range m {
		keys = append(keys, key)
	}

	return
}
