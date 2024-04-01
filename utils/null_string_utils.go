package utils

import "strings"

// TagBytesToString converts a byte array to a string.
func BytesToStringTrimNull(b []byte) string {
	s := string(b)
	return strings.Trim(s, "\x00")
}

// remove the null bytes
func TrimNull(s string) string {
	return strings.TrimRight(s, "\x00")
}
