package ucl

import (
	"fmt"
)

%%{
	machine common;
	alphtype rune;

	action error {
		return /*nil, */fmt.Errorf("parse error at byte %d (state=%d)", fpc, cs)
	}

	action done { fhold; fbreak; }
	
	ws = [ \t\r\n];

	false = "false";
	true = "true";
	nullval = "null";

	unescaped = (0x20..0x21 | 0x23..0x5B | 0x5D..0x10FFFF);
	char = (unescaped | "\\" . ([\"\\/bfnrt] | "u" . [0-9a-fA-F]{4}));
	string = ("\"" . char** . "\"");

	int = "0" | ([1-9] . [0-9]*);
	number = "-"? . int . ("." . [0-9]+)? . ([eE] . [\+\-]? . [0-9]+)? (^[0-9eE\-\+.] @done);
}%%

%%{
	machine object;
	include common;
	
	action parse_value {
		newp, err := parse_value(data, fpc, pe);
		if err != nil { return -1, err };
		fexec newp;
	}

	member = (string . ws* . ":" . ws* . (^ws >parse_value));

	object_content = (member . (ws* . "," . ws* . member)*);
	
	main := (ws* . "{" . ws* . object_content? . ws* . ("}" %*done));
	
	write data;
}%%

func parse_object(data []rune, p int, pe int) (int, error) {
	var (
		cs int
		eof = pe
	)
	_ = eof

%% write init;
%% write exec;

	if cs >= object_first_final {
		return p, nil
	}
	return -1, fmt.Errorf("[object] wat p=%d cs=%d", p, cs)
}

%%{
	machine array;
	include common;
	
	action parse_value {
		newp, err := parse_value(data, fpc, pe);
		if err != nil { return -1, err };
		fexec newp;
	}

	value = ^(ws | "]") >parse_value;

	array_content = (value . (ws* . "," . ws* . value)*);
	
	main := (ws* . "[" . ws* . array_content? . ws* . ("]" %*done));
	
	write data;
}%%

func parse_array(data []rune, p int, pe int) (int, error) {
	var (
		cs int
		eof = pe
	)
	_ = eof

%% write init;
%% write exec;

	if cs >= array_first_final {
		return p, nil
	}
	return -1, fmt.Errorf("[array] wat p=%d cs=%d", p, cs)
}

%%{
	machine value;
	include common;

	action parse_value_error {
		return -1, fmt.Errorf("parse error at byte %d (state=%d)", fpc, cs)
	}
	
	action parse_object {
		newp, err := parse_object(data, fpc, pe);
		if err != nil { return -1, err };
		fexec newp;
	}

	action parse_array {
		newp, err := parse_array(data, fpc, pe);
		if err != nil { return -1, err };
		fexec newp;
	}
	
	array = ws* . ("[" >parse_array);
	object = ws* . ("{" >parse_object);
	
	main := (false | true | nullval | object | array | number | string) %*done $!parse_value_error;
	
	write data;
}%%

func parse_value(data []rune, p int, pe int) (int, error) {
	var (
		cs int
		eof = pe
	)

%% write init;
%% write exec;
	if cs >= value_first_final {
		return p, nil
	}
	return -1, fmt.Errorf("wat p=%d cs=%d", p, cs)
}

%%{
	machine document;
	include common;

	action parse_object {
		newp, err := parse_object(data, fpc, pe);
		if err != nil { return err };
		fexec newp;
	}

	action parse_array {
		newp, err := parse_array(data, fpc, pe);
		if err != nil { return err };
		fexec newp;
	}
	
	array = ws* . ("[" >parse_array);	
	object = ws* . ("{" >parse_object);

	document = (object | array);
	
	main := document $!error;
	
	write data;
}%%

func parse_json(data []rune) (/*Document,*/ error) {
	var (
		cs int
		p int
		pe int = len(data)
		eof int = len(data)
		top int
		stack []int
	)
	_ = top
	_ = stack
	var (
//		ret = Document{}
	)

%% write init;
%% write exec;

	return /*ret, */nil
}