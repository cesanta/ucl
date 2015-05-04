package ucl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//go:generate ragel -Z ucl.rl

type Value interface {
	String() string
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

type Bool struct {
	Value bool
}

func (v Bool) String() string {
	if v.Value {
		return "true"
	}
	return "false"
}

type Number struct {
	Value float64
}

func (v Number) String() string {
	return fmt.Sprintf("%g", v.Value)
}

type String struct {
	Value string
}

func (v String) String() string {
	// TODO(imax): this does not necessarily emits valid JSON.
	return fmt.Sprintf("%q", v.Value)
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

type Key struct {
	Value string
}

func (v Key) String() string {
	// TODO(imax): this does not necessarily emits valid JSON.
	return fmt.Sprintf("%q", v.Value)
}

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
