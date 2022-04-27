package str

import (
	"strings"
	"unicode"
)

func ToSnackCase(s string, upper bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	adjustCase := unicode.ToLower
	if upper {
		adjustCase = unicode.ToUpper
	}

	var (
		builder strings.Builder
		prev    rune
		curr    rune
	)
	for _, next := range s {
		switch {
		case !unicode.IsLetter(curr) && !unicode.IsDigit(curr):
			if unicode.IsLetter(prev) || unicode.IsDigit(prev) {
				builder.WriteRune('_')
			}
		case unicode.IsUpper(curr):
			if unicode.IsLower(prev) || unicode.IsDigit(prev) || (unicode.IsUpper(prev) && unicode.IsLower(next)) {
				builder.WriteRune('_')
			}
			builder.WriteRune(adjustCase(curr))
		default:
			builder.WriteRune(adjustCase(curr))
		}

		prev = curr
		curr = next
	}

	if unicode.IsUpper(curr) && unicode.IsLower(prev) && prev != 0 {
		builder.WriteRune('_')
	}
	builder.WriteRune(adjustCase(curr))

	return builder.String()
}
