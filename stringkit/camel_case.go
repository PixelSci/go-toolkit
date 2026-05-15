package stringkit

import "unicode"

// CamelCase converts a string to camelCase.
// It handles snake_case, kebab-case, space-separated, PascalCase, and ALL_CAPS inputs.
// Separators (underscore, hyphen, space) and uppercase-following-lowercase are treated
// as word boundaries.
//
//	CamelCase("hello_world")   → "helloWorld"
//	CamelCase("hello-world")   → "helloWorld"
//	CamelCase("HelloWorld")    → "helloWorld"
//	CamelCase("HELLO_WORLD")   → "helloWorld"
func CamelCase(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	var result []rune
	nextUpper := false

	for i, r := range runes {
		if r == '_' || r == '-' || r == ' ' {
			if len(result) > 0 {
				nextUpper = true
			}
			continue
		}

		if nextUpper {
			result = append(result, unicode.ToUpper(r))
			nextUpper = false
			continue
		}

		if i == 0 {
			result = append(result, unicode.ToLower(r))
			continue
		}

		// Uppercase following a lowercase letter marks a new word (PascalCase boundary).
		if unicode.IsUpper(r) && unicode.IsLower(runes[i-1]) {
			result = append(result, r)
			continue
		}

		if unicode.IsUpper(r) {
			result = append(result, unicode.ToLower(r))
			continue
		}

		result = append(result, r)
	}

	return string(result)
}
