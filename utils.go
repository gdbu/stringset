package stringset

import (
	"fmt"
)

const (
	stateStart uint8 = iota
	statePreQuote
	stateValue
	stateEscape
	statePostQuote
)

func forEachKey(bs []byte, fn func(key string)) (err error) {
	if len(bs) == 0 {
		return
	}

	var state uint8
	keyBuf := make([]byte, 0, 32)
	for i := 0; i < len(bs); i++ {
		char := bs[i]
		switch state {
		case stateStart:
			switch char {
			case '[':
				state = statePreQuote
			default:
				return fmt.Errorf("invalid character, expected <[>, received <%s>", string(char))
			}
		case statePreQuote:
			switch char {
			case '"':
				state = stateValue
			case ' ', '\t':
			default:
				return fmt.Errorf("invalid character, expected <\">, < >, or <\t>, received <%s>", string(char))
			}

		case stateValue:
			switch char {
			case '\\':
				state = stateEscape
			case '"':
				state = statePostQuote
			default:
				keyBuf = append(keyBuf, char)
			}

		case stateEscape:
			keyBuf = append(keyBuf, char)
			state = stateValue

		case statePostQuote:
			fn(string(keyBuf))
			keyBuf = keyBuf[:0]

			switch char {
			case ',':
				state = statePreQuote
			case ']':
				return
			default:
				return fmt.Errorf("invalid character, expected <,> or <]>, received <%s>", string(char))
			}
		}
	}

	return
}
