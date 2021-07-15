package stringset

import (
	"errors"
	"testing"
)

func Test_forEachKey(t *testing.T) {
	type testcase struct {
		str       string
		wanted    []string
		wantedErr error
	}

	tcs := []testcase{
		{
			str:       `["foo","bar","baz"]`,
			wanted:    []string{"foo", "bar", "baz"},
			wantedErr: nil,
		},
		{
			str:       `["foo","bar","baz",]`,
			wanted:    []string{"foo", "bar", "baz"},
			wantedErr: errors.New("invalid character, expected <\">, < >, or <\t>, received <]>"),
		},

		{
			str:       `{"foo":1,"bar":2,"baz":3}`,
			wanted:    []string{},
			wantedErr: errors.New("invalid character, expected <[>, received <{>"),
		},
	}

	for _, tc := range tcs {
		var count int
		err := forEachKey([]byte(tc.str), func(key string) {
			if count+1 > len(tc.wanted) {
				t.Fatalf("exceeded wanted count, expected %d entries <%s>", len(tc.wanted), tc.str)
			}

			if tc.wanted[count] != key {
				t.Fatalf("invalid key, expected <%s> and received <%s>", tc.wanted[count], key)
			}

			count++
		})

		if !checkErr(tc.wantedErr, err) {
			t.Fatalf("invalid error, expected\n<%v> and received\n<%v>", tc.wantedErr, err)
		}

		if count != len(tc.wanted) {
			t.Fatalf("invalid number of entries, expected %d and received %d", len(tc.wanted), count)
		}
	}
}

func checkErr(a, b error) bool {
	switch {
	case a == nil && b == nil:
		return true
	case a == nil && b != nil:
		return false
	case a != nil && b == nil:
		return false

	default:
		return a.Error() == b.Error()
	}
}
