//this package encodes and decodes in ROT13
package gorot

import "strings"

func indexPosition(s string, arr []string) int {
	for i, c := range arr {
		if c == s {
			return i
		}
	}

	return -1
}

// Encode takes a plain text message and returns the rot13 encoded result.
func Encode(message string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	alphaLen := len(alphabet)

	var encoded []string

	for _, c := range message {
		strC := string(c)
		if strC == " " {
			encoded = append(encoded, " ")

			continue
		}

		isUpper := false
		if strings.ToLower(strC) != strC {
			isUpper = true
		}

		i := indexPosition(strings.ToLower(strC), alphabet)
		if i == -1 {
			encoded = append(encoded, strC)

			continue
		}

		var pos int
		pos = i + 13

		if pos >= alphaLen {
			pos = (i + 13) - alphaLen
		}

		if isUpper {
			encoded = append(encoded, strings.ToUpper(alphabet[pos]))

			continue
		}

		encoded = append(encoded, alphabet[pos])
	}

	return strings.Join(encoded, "")
}

// Decode takes a rot13 encoded message and returns the unencoded result.
func Decode(message string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	alphaLen := len(alphabet)

	var decoded []string

	for _, c := range message {
		strC := string(c)
		if strC == " " {
			decoded = append(decoded, " ")

			continue
		}

		isUpper := false
		if strings.ToLower(strC) != strC {
			isUpper = true
		}

		i := indexPosition(strings.ToLower(strC), alphabet)
		if i == -1 {
			decoded = append(decoded, strC)

			continue
		}

		var pos int
		pos = i - 13

		if pos < 0 {
			pos = alphaLen + (i - 13)
		}

		if isUpper {
			decoded = append(decoded, strings.ToUpper(alphabet[pos]))

			continue
		}

		decoded = append(decoded, alphabet[pos])
	}

	return strings.Join(decoded, "")
}
