// Code generated by "stringer -type=OpCode"; DO NOT EDIT.

package laks

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OP_PUSH-0]
	_ = x[OP_ADD-1]
	_ = x[OP_MULT-2]
	_ = x[OP_PRINT-3]
}

const _OpCode_name = "OP_PUSHOP_ADDOP_MULTOP_PRINT"

var _OpCode_index = [...]uint8{0, 7, 13, 20, 28}

func (i OpCode) String() string {
	if i >= OpCode(len(_OpCode_index)-1) {
		return "OpCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OpCode_name[_OpCode_index[i]:_OpCode_index[i+1]]
}
