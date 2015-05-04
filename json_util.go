package ucl

import (
	"strings"
)

// TODO(imax): implement proper escaping/unescaping.

func json_escape(s string) string {
	return strings.Replace(s, "\"", "\\\"", -1)
}

func json_unescape(s string) string {
	return strings.Replace(s, "\\\"", "\"", -1)
}
