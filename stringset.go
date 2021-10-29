package stringset

type StringSet interface {
	Set(key string)
	SetMany(keys []string)
	Unset(key string)
	UnsetMany(keys []string)
	Has(key string) (has bool)
	HasAll(keys []string) (has bool)
	HasOne(keys []string) (has bool)
	Slice() (keys []string)
	IsMatch(in StringSet) (isMatch bool)
	SetAsString(value string) (err error)
	Len() (n int)
}
