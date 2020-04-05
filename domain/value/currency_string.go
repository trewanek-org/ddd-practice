// Code generated by "stringer -type Currency ./currency.go"; DO NOT EDIT.

package value

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[JPY-1]
	_ = x[USD-2]
	_ = x[EUR-3]
}

const _Currency_name = "JPYUSDEUR"

var _Currency_index = [...]uint8{0, 3, 6, 9}

func (i Currency) String() string {
	i -= 1
	if i < 0 || i >= Currency(len(_Currency_index)-1) {
		return "Currency(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Currency_name[_Currency_index[i]:_Currency_index[i+1]]
}