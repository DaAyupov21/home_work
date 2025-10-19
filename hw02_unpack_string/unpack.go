package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var (
		b       strings.Builder
		prev    rune
		hasPrev bool
		escape  bool
	)
	b.Grow(len(s) * 2)

	writeN := func(r rune, n int) {
		for i := 0; i < n; i++ {
			b.WriteRune(r)
		}
	}

	for i := 0; i < len(s); {
		r, sz := utf8.DecodeRuneInString(s[i:])
		i += sz

		switch {
		case escape:
			// предыдущий символ можно окончательно вывести
			if hasPrev {
				b.WriteRune(prev)
			}
			prev, hasPrev = r, true
			escape = false

		case r == '\\':
			escape = true

		case r >= '0' && r <= '9':
			if !hasPrev {
				return "", ErrInvalidString
			}
			rep := int(r - '0')
			if rep == 0 {
				// «стираем» предыдущий символ
				hasPrev = false
				continue
			}
			// пишем ровно rep штук и сбрасываем «висящий»
			writeN(prev, rep)
			hasPrev = false

		default:
			// обычный символ — дописываем предыдущий, подвешиваем текущий
			if hasPrev {
				b.WriteRune(prev)
			}
			prev, hasPrev = r, true
		}
	}

	if escape {
		return "", ErrInvalidString // строка закончилась на '\'
	}
	if hasPrev {
		b.WriteRune(prev)
	}
	return b.String(), nil
}
