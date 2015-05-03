
//line ucl.rl:1
package ucl

import (
	"fmt"
)


//line ucl.rl:29



//line ucl.go:15
var _object_actions []byte = []byte{
	0, 1, 0, 1, 1, 
}

var _object_key_offsets []byte = []byte{
	0, 0, 5, 11, 15, 20, 24, 30, 
	35, 44, 50, 56, 62, 68, 
}

var _object_trans_keys []int32 = []int32{
	13, 32, 123, 9, 10, 13, 32, 34, 
	125, 9, 10, 34, 92, 32, 1114111, 13, 
	32, 58, 9, 10, 13, 32, 9, 10, 
	13, 32, 44, 125, 9, 10, 13, 32, 
	34, 9, 10, 34, 47, 92, 98, 102, 
	110, 114, 116, 117, 48, 57, 65, 70, 
	97, 102, 48, 57, 65, 70, 97, 102, 
	48, 57, 65, 70, 97, 102, 48, 57, 
	65, 70, 97, 102, 
}

var _object_single_lengths []byte = []byte{
	0, 3, 4, 2, 3, 2, 4, 3, 
	9, 0, 0, 0, 0, 0, 
}

var _object_range_lengths []byte = []byte{
	0, 1, 1, 1, 1, 1, 1, 1, 
	0, 3, 3, 3, 3, 0, 
}

var _object_index_offsets []byte = []byte{
	0, 0, 5, 11, 15, 20, 24, 30, 
	35, 45, 49, 53, 57, 61, 
}

var _object_indicies []byte = []byte{
	0, 0, 2, 0, 1, 2, 2, 3, 
	4, 2, 1, 5, 6, 3, 1, 5, 
	5, 7, 5, 1, 7, 7, 7, 8, 
	9, 9, 10, 4, 9, 1, 10, 10, 
	3, 10, 1, 3, 3, 3, 3, 3, 
	3, 3, 3, 11, 1, 12, 12, 12, 
	1, 13, 13, 13, 1, 14, 14, 14, 
	1, 3, 3, 3, 1, 1, 
}

var _object_trans_targs []byte = []byte{
	1, 0, 2, 3, 13, 4, 8, 5, 
	6, 6, 7, 9, 10, 11, 12, 
}

var _object_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	3, 0, 0, 0, 0, 0, 0, 
}

var _object_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 
}

const object_start int = 1
const object_first_final int = 13
const object_error int = 0

const object_en_main int = 1


//line ucl.rl:48


func parse_object(data []rune, p int, pe int) (int, error) {
	var (
		cs int
		eof = pe
	)
	_ = eof


//line ucl.go:96
	{
	cs = object_start
	}

//line ucl.rl:58

//line ucl.go:103
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_object_from_state_actions[cs])
	_nacts = uint(_object_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _object_actions[_acts - 1] {
		case 0:
//line ucl.rl:15
 p--
 p++; goto _out
 
//line ucl.go:127
		}
	}

	_keys = int(_object_key_offsets[cs])
	_trans = int(_object_index_offsets[cs])

	_klen = int(_object_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _object_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _object_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_object_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _object_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _object_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_object_indicies[_trans])
	cs = int(_object_trans_targs[_trans])

	if _object_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_object_trans_actions[_trans])
	_nacts = uint(_object_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _object_actions[_acts-1] {
		case 1:
//line ucl.rl:35

		newp, err := parse_value(data, p, pe);
		if err != nil { return -1, err };
		p = ( newp) - 1

	
//line ucl.go:204
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line ucl.rl:59

	if cs >= object_first_final {
		return p, nil
	}
	return -1, fmt.Errorf("[object] wat p=%d cs=%d", p, cs)
}


//line ucl.go:229
var _array_actions []byte = []byte{
	0, 1, 0, 1, 1, 
}

var _array_key_offsets []byte = []byte{
	0, 0, 5, 10, 16, 21, 
}

var _array_trans_keys []int32 = []int32{
	13, 32, 91, 9, 10, 13, 32, 93, 
	9, 10, 13, 32, 44, 93, 9, 10, 
	13, 32, 93, 9, 10, 
}

var _array_single_lengths []byte = []byte{
	0, 3, 3, 4, 3, 0, 
}

var _array_range_lengths []byte = []byte{
	0, 1, 1, 1, 1, 0, 
}

var _array_index_offsets []byte = []byte{
	0, 0, 5, 10, 16, 21, 
}

var _array_indicies []byte = []byte{
	0, 0, 2, 0, 1, 2, 2, 4, 
	2, 3, 5, 5, 6, 4, 5, 1, 
	6, 6, 1, 6, 3, 1, 
}

var _array_trans_targs []byte = []byte{
	1, 0, 2, 3, 5, 3, 4, 
}

var _array_trans_actions []byte = []byte{
	0, 0, 0, 3, 0, 0, 0, 
}

var _array_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 1, 
}

const array_start int = 1
const array_first_final int = 5
const array_error int = 0

const array_en_main int = 1


//line ucl.rl:83


func parse_array(data []rune, p int, pe int) (int, error) {
	var (
		cs int
		eof = pe
	)
	_ = eof


//line ucl.go:292
	{
	cs = array_start
	}

//line ucl.rl:93

//line ucl.go:299
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_array_from_state_actions[cs])
	_nacts = uint(_array_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _array_actions[_acts - 1] {
		case 0:
//line ucl.rl:15
 p--
 p++; goto _out
 
//line ucl.go:323
		}
	}

	_keys = int(_array_key_offsets[cs])
	_trans = int(_array_index_offsets[cs])

	_klen = int(_array_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _array_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _array_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_array_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _array_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _array_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_array_indicies[_trans])
	cs = int(_array_trans_targs[_trans])

	if _array_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_array_trans_actions[_trans])
	_nacts = uint(_array_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _array_actions[_acts-1] {
		case 1:
//line ucl.rl:70

		newp, err := parse_value(data, p, pe);
		if err != nil { return -1, err };
		p = ( newp) - 1

	
//line ucl.go:400
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line ucl.rl:94

	if cs >= array_first_final {
		return p, nil
	}
	return -1, fmt.Errorf("[array] wat p=%d cs=%d", p, cs)
}


//line ucl.go:425
var _value_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
}

var _value_key_offsets []byte = []byte{
	0, 0, 14, 20, 24, 33, 39, 45, 
	51, 57, 60, 67, 69, 76, 80, 82, 
	89, 96, 97, 98, 99, 100, 101, 102, 
	103, 104, 105, 
}

var _value_trans_keys []int32 = []int32{
	13, 32, 34, 45, 48, 91, 102, 110, 
	116, 123, 9, 10, 49, 57, 13, 32, 
	91, 123, 9, 10, 34, 92, 32, 1114111, 
	34, 47, 92, 98, 102, 110, 114, 116, 
	117, 48, 57, 65, 70, 97, 102, 48, 
	57, 65, 70, 97, 102, 48, 57, 65, 
	70, 97, 102, 48, 57, 65, 70, 97, 
	102, 48, 49, 57, 43, 45, 46, 69, 
	101, 48, 57, 48, 57, 43, 69, 101, 
	45, 46, 48, 57, 43, 45, 48, 57, 
	48, 57, 43, 69, 101, 45, 46, 48, 
	57, 43, 45, 46, 69, 101, 48, 57, 
	97, 108, 115, 101, 117, 108, 108, 114, 
	117, 
}

var _value_single_lengths []byte = []byte{
	0, 10, 4, 2, 9, 0, 0, 0, 
	0, 1, 5, 0, 3, 2, 0, 3, 
	5, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 0, 
}

var _value_range_lengths []byte = []byte{
	0, 2, 1, 1, 0, 3, 3, 3, 
	3, 1, 1, 1, 2, 1, 1, 2, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 
}

var _value_index_offsets []byte = []byte{
	0, 0, 13, 19, 23, 33, 37, 41, 
	45, 49, 52, 59, 61, 67, 71, 73, 
	79, 86, 88, 90, 92, 94, 96, 98, 
	100, 102, 104, 
}

var _value_indicies []byte = []byte{
	1, 1, 2, 3, 4, 6, 7, 8, 
	9, 10, 1, 5, 0, 1, 1, 6, 
	10, 1, 0, 11, 12, 2, 0, 2, 
	2, 2, 2, 2, 2, 2, 2, 13, 
	0, 14, 14, 14, 0, 15, 15, 15, 
	0, 16, 16, 16, 0, 2, 2, 2, 
	0, 4, 5, 0, 0, 0, 18, 19, 
	19, 0, 17, 20, 0, 0, 19, 19, 
	0, 20, 17, 21, 21, 22, 0, 22, 
	0, 0, 0, 0, 0, 22, 17, 0, 
	0, 18, 19, 19, 5, 17, 23, 0, 
	24, 0, 25, 0, 11, 0, 26, 0, 
	27, 0, 11, 0, 28, 0, 25, 0, 
	0, 
}

var _value_trans_targs []byte = []byte{
	0, 2, 3, 9, 10, 16, 26, 17, 
	21, 24, 26, 26, 4, 5, 6, 7, 
	8, 26, 11, 13, 12, 14, 15, 18, 
	19, 20, 22, 23, 25, 
}

var _value_trans_actions []byte = []byte{
	3, 0, 0, 0, 0, 0, 7, 0, 
	0, 0, 5, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _value_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 
}

var _value_eof_actions []byte = []byte{
	0, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 0, 
}

const value_start int = 1
const value_first_final int = 26
const value_error int = 0

const value_en_main int = 1


//line ucl.rl:127


func parse_value(data []rune, p int, pe int) (int, error) {
	var (
		cs int
		eof = pe
	)


//line ucl.go:537
	{
	cs = value_start
	}

//line ucl.rl:136

//line ucl.go:544
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_value_from_state_actions[cs])
	_nacts = uint(_value_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _value_actions[_acts - 1] {
		case 0:
//line ucl.rl:15
 p--
 p++; goto _out
 
//line ucl.go:568
		}
	}

	_keys = int(_value_key_offsets[cs])
	_trans = int(_value_index_offsets[cs])

	_klen = int(_value_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _value_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _value_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_value_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _value_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _value_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_value_indicies[_trans])
	cs = int(_value_trans_targs[_trans])

	if _value_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_value_trans_actions[_trans])
	_nacts = uint(_value_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _value_actions[_acts-1] {
		case 0:
//line ucl.rl:15
 p--
 p++; goto _out
 
		case 1:
//line ucl.rl:105

		return -1, fmt.Errorf("parse error at byte %d (state=%d)", p, cs)
	
		case 2:
//line ucl.rl:109

		newp, err := parse_object(data, p, pe);
		if err != nil { return -1, err };
		p = ( newp) - 1

	
		case 3:
//line ucl.rl:115

		newp, err := parse_array(data, p, pe);
		if err != nil { return -1, err };
		p = ( newp) - 1

	
//line ucl.go:663
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _value_eof_actions[cs]
		__nacts := uint(_value_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _value_actions[__acts-1] {
			case 1:
//line ucl.rl:105

		return -1, fmt.Errorf("parse error at byte %d (state=%d)", p, cs)
	
//line ucl.go:687
			}
		}
	}

	_out: {}
	}

//line ucl.rl:137
	if cs >= value_first_final {
		return p, nil
	}
	return -1, fmt.Errorf("wat p=%d cs=%d", p, cs)
}


//line ucl.go:703
var _document_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 
}

var _document_key_offsets []byte = []byte{
	0, 0, 6, 
}

var _document_trans_keys []int32 = []int32{
	13, 32, 91, 123, 9, 10, 
}

var _document_single_lengths []byte = []byte{
	0, 4, 0, 
}

var _document_range_lengths []byte = []byte{
	0, 1, 0, 
}

var _document_index_offsets []byte = []byte{
	0, 0, 6, 
}

var _document_trans_targs []byte = []byte{
	1, 1, 2, 2, 1, 0, 0, 
}

var _document_trans_actions []byte = []byte{
	0, 0, 5, 3, 0, 1, 1, 
}

var _document_eof_actions []byte = []byte{
	0, 1, 0, 
}

const document_start int = 1
const document_first_final int = 2
const document_error int = 0

const document_en_main int = 1


//line ucl.rl:167


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


//line ucl.go:766
	{
	cs = document_start
	}

//line ucl.rl:185

//line ucl.go:773
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_document_key_offsets[cs])
	_trans = int(_document_index_offsets[cs])

	_klen = int(_document_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _document_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _document_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_document_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _document_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _document_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_document_trans_targs[_trans])

	if _document_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_document_trans_actions[_trans])
	_nacts = uint(_document_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _document_actions[_acts-1] {
		case 0:
//line ucl.rl:11

		return /*nil, */fmt.Errorf("parse error at byte %d (state=%d)", p, cs)
	
		case 1:
//line ucl.rl:147

		newp, err := parse_object(data, p, pe);
		if err != nil { return err };
		p = ( newp) - 1

	
		case 2:
//line ucl.rl:153

		newp, err := parse_array(data, p, pe);
		if err != nil { return err };
		p = ( newp) - 1

	
//line ucl.go:872
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _document_eof_actions[cs]
		__nacts := uint(_document_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _document_actions[__acts-1] {
			case 0:
//line ucl.rl:11

		return /*nil, */fmt.Errorf("parse error at byte %d (state=%d)", p, cs)
	
//line ucl.go:896
			}
		}
	}

	_out: {}
	}

//line ucl.rl:186

	return /*ret, */nil
}