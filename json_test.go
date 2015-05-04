package ucl

import (
	"testing"
)

func TestParseJSON(t *testing.T) {
	good := []string{
		`{}`,
		`[]`,
		`[true, false, null]`,
		`[123]`,
		`{"1":"2"}`,
		`{"a":123}`,
		`{"a":1,"b":"321"}`,
		`{"a":123.45}`,
		`{"a":123.45e67}`,
		`{"a":-123}`,
		`{"a":-123.45}`,
		`{"a":-123.45e-67}`,
		`{"a":["b"]}`,
		`{"a":{"b":{"c":123}}}`,
		`[[[123,"1 ", "1", 123 , {}]]]`,
	}
	for _, s := range good {
		if v, err := parse([]rune(s)); err != nil {
			t.Errorf("Failed to parse '%s': %s", s, err)
		} else {
			t.Logf("'%s' -> %s", s, v)
		}
	}
	bad := []string{
		`""`,
		`["""]`,
		`{}{}`,
		`{{}}`,
		`{[]:{}}`,
		`{{{{{`,
		`}{`,
		`][`,
		`"a"a"`,
	}
	for _, s := range bad {
		_, err := parse([]rune(s))
		if err == nil {
			t.Errorf("Parse succeeded on invalid JSON document '%s'", s)
		} else {
			t.Logf("%s: %s", s, err)
		}
	}
}
