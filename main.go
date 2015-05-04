package ucl

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

//go:generate ragel -Z ucl.rl

type FormatConfig struct {
	Indent                   string `json:",omitempty"`
	MultilineObjectThreshold int    `json:",omitempty"`
	MultilineArrayThreshold  int    `json:",omitempty"`
	PreserveObjectKeysOrder  bool   `json:",omitempty"`
}

type Value interface {
	String() string
	format(indent string, config *FormatConfig) string
}

func Format(v Value, c *FormatConfig, w io.Writer) error {
	if c == nil {
		c = &FormatConfig{} // keep this empty, zero values for options should mean default format.
	}
	_, err := w.Write([]byte(v.format("", c)))
	return err
}

func Parse(r io.Reader) (Value, error) {
	rr := bufio.NewReader(r)
	data := []rune{}
	for {
		c, _, err := rr.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, c)
	}
	return parse(data)
}

func parse(data []rune) (Value, error) {
	v, _, err := parse_json(data)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type Null struct{}

func (Null) String() string {
	return "null"
}

func (Null) format(indent string, config *FormatConfig) string {
	return "null"
}

type Bool struct {
	Value bool
}

func (v Bool) String() string {
	if v.Value {
		return "true"
	}
	return "false"
}

func (v Bool) format(indent string, config *FormatConfig) string {
	return v.String()
}

type Number struct {
	Value float64
}

func (v Number) String() string {
	return fmt.Sprintf("%g", v.Value)
}

func (v Number) format(indent string, config *FormatConfig) string {
	return v.String()
}

type String struct {
	Value string
}

func (v String) String() string {
	// TODO(imax): this does not necessarily emits valid JSON.
	return fmt.Sprintf("%q", v.Value)
}

func (v String) format(indent string, config *FormatConfig) string {
	return v.String()
}

type Array struct {
	Value []Value
}

func (v Array) String() string {
	t := make([]string, len(v.Value))
	for i, item := range v.Value {
		t[i] = item.String()
	}
	return "[" + strings.Join(t, ",") + "]"
}

func (v Array) format(indent string, config *FormatConfig) string {
	if len(v.Value) == 0 {
		return "[]"
	}
	newIndent := config.Indent
	if newIndent == "" {
		newIndent = "  "
	}
	items := make([]string, len(v.Value))
	for i, item := range v.Value {
		items[i] = item.format(indent+newIndent, config)
	}
	shortFormatLen := 2                    // brackets
	shortFormatLen += (len(items) - 1) * 2 // ", " between items
	for _, item := range items {
		if strings.IndexRune(item, '\n') >= 0 {
			// One of the items spans multiple lines, bail out.
			shortFormatLen = config.MultilineArrayThreshold + 1
			break
		}
		shortFormatLen += len(item)
		if shortFormatLen > config.MultilineArrayThreshold {
			break
		}
	}
	if shortFormatLen <= config.MultilineArrayThreshold {
		return "[" + strings.Join(items, ", ") + "]"
	}
	r := "[\n"
	for i := 0; i < len(items)-1; i++ {
		r += indent + newIndent + items[i] + ",\n"
	}
	r += indent + newIndent + items[len(items)-1] + "\n"
	r += indent + "]"
	return r
}

type Key struct {
	Value string
	Index int
}

func (v Key) String() string {
	// TODO(imax): this does not necessarily emits valid JSON.
	return fmt.Sprintf("%q", v.Value)
}

func (v Key) format(indent string, config *FormatConfig) string {
	return v.String()
}

type byValue []Key

func (s byValue) Len() int           { return len(s) }
func (s byValue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byValue) Less(i, j int) bool { return s[i].Value < s[j].Value }

type byIndex []Key

func (s byIndex) Len() int           { return len(s) }
func (s byIndex) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byIndex) Less(i, j int) bool { return s[i].Index < s[j].Index }

type Object struct {
	Value map[Key]Value
}

func (v Object) String() string {
	t := make([]string, 0, len(v.Value))
	for key, item := range v.Value {
		t = append(t, key.String()+":"+item.String())
	}
	return "{" + strings.Join(t, ",") + "}"
}

func (v Object) format(indent string, config *FormatConfig) string {
	if len(v.Value) == 0 {
		return "{}"
	}
	newIndent := config.Indent
	if newIndent == "" {
		newIndent = "  "
	}
	// Make sure that order of properties is stable.
	keys := make([]Key, 0, len(v.Value))
	for k := range v.Value {
		keys = append(keys, k)
	}
	if config.PreserveObjectKeysOrder {
		sort.Sort(byIndex(keys))
	} else {
		sort.Sort(byValue(keys))
	}

	items := make([]string, len(keys))
	for i, k := range keys {
		items[i] = k.format(indent+newIndent, config) + ": " + v.Value[k].format(indent+newIndent, config)
	}
	shortFormatLen := 2                    // brackets
	shortFormatLen += (len(items) - 1) * 2 // ", " between items
	for _, item := range items {
		if strings.IndexRune(item, '\n') >= 0 {
			// One of the items spans multiple lines, bail out.
			shortFormatLen = config.MultilineObjectThreshold + 1
			break
		}
		shortFormatLen += len(item)
		if shortFormatLen > config.MultilineObjectThreshold {
			break
		}
	}
	if shortFormatLen <= config.MultilineObjectThreshold {
		return "{" + strings.Join(items, ", ") + "}"
	}

	r := "{\n"
	for i := 0; i < len(keys)-1; i++ {
		r += indent + newIndent + keys[i].format(indent+newIndent, config) + ": " + v.Value[keys[i]].format(indent+newIndent, config) + ",\n"
	}
	r += indent + newIndent + keys[len(keys)-1].format(indent+newIndent, config) + ": " + v.Value[keys[len(keys)-1]].format(indent+newIndent, config) + "\n"
	r += indent + "}"
	return r
}
