package ucl

import (
	"bytes"
	"fmt"
)

/* TODO(imax): code below does not take into account escaping of
 * non-BMP characters as specified by RFC 4627.
 *
 *   To escape an extended character that is not in the Basic Multilingual
 *   Plane, the character is represented as a twelve-character sequence,
 *   encoding the UTF-16 surrogate pair.  So, for example, a string
 *   containing only the G clef character (U+1D11E) may be represented as
 *   "\uD834\uDD1E".
 *
 * Such surrogate pairs will be unescaped as 2 adjacent UTF-8 sequences.
 */

func jsonEscape(s string) string {
	r := bytes.NewBuffer(nil)
	for _, c := range s {
		switch c {
		case '"':
			r.WriteString(`\"`) // err is always nil
		case '\\':
			r.WriteString(`\\`) // err is always nil
		case '\b':
			r.WriteString(`\b`) // err is always nil
		case '\f':
			r.WriteString(`\f`) // err is always nil
		case '\n':
			r.WriteString(`\n`) // err is always nil
		case '\r':
			r.WriteString(`\r`) // err is always nil
		case '\t':
			r.WriteString(`\t`) // err is always nil
		default:
			r.WriteRune(c) // err is always nil
		}
	}
	return r.String()
}

func isHexDigit(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

func hexDigitValue(c rune) rune {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10
	}
	return '\uFFFD'
}

func jsonUnescape(s string) (string, error) {
	r := bytes.NewBuffer(nil)
	start := 0
	const (
		RecordStart    = iota // Beginning of the string or right after the escape sequence.
		Regular               // After a regular unescaped character.
		AfterBackslash        // After a backslash.
		AfterU                // After \u, a separate counter is used to eat exactly 4 next characters.
	)
	state := RecordStart
	var ucount uint
	var u rune
	// Iteration over a string interpretes it as UTF-8 and produces Unicode
	// runes. i is the index of the first byte of the rune, c is the rune.
	for i, c := range s {
		switch state {
		case RecordStart:
			switch c {
			case '\\':
				state = AfterBackslash
			default:
				start = i
				state = Regular
			}
		case Regular:
			switch c {
			case '\\':
				r.WriteString(s[start:i]) // err is always nil
				state = AfterBackslash
			}
		case AfterBackslash:
			switch c {
			case '"':
				r.WriteString("\"") // err is always nil
				state = RecordStart
			case '\\':
				r.WriteString("\\") // err is always nil
				state = RecordStart
			case '/':
				r.WriteString("/") // err is always nil
				state = RecordStart
			case 'b':
				r.WriteString("\b") // err is always nil
				state = RecordStart
			case 'f':
				r.WriteString("\f") // err is always nil
				state = RecordStart
			case 'n':
				r.WriteString("\n") // err is always nil
				state = RecordStart
			case 'r':
				r.WriteString("\r") // err is always nil
				state = RecordStart
			case 't':
				r.WriteString("\t") // err is always nil
				state = RecordStart
			case 'u':
				ucount = 0
				u = 0
				state = AfterU
			default:
				return "", fmt.Errorf("invalid escape sequence %q at %d", c, i)
			}
		case AfterU:
			if !isHexDigit(c) {
				return "", fmt.Errorf("invalid hex digit %q at %d", c, i)
			}
			v := hexDigitValue(c)
			u |= v << (4 * (3 - ucount))
			ucount++
			if ucount == 4 {
				r.WriteRune(u) // err is always nil
				state = RecordStart
			}
		}
	}
	if state == Regular {
		r.WriteString(s[start:len(s)]) // err is always nil
	}
	if state != Regular && state != RecordStart {
		return "", fmt.Errorf("incomplete escape sequence at the end of the string")
	}
	return r.String(), nil
}
