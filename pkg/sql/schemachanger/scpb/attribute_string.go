// Code generated by "stringer -type=Attribute -trimprefix=Attribute"; DO NOT EDIT.

package scpb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AttributeType-1]
	_ = x[AttributeDescID-2]
	_ = x[AttributeDepID-3]
	_ = x[AttributeColumnID-4]
	_ = x[AttributeElementName-5]
	_ = x[AttributeIndexID-6]
}

const _Attribute_name = "TypeDescIDDepIDColumnIDElementNameIndexID"

var _Attribute_index = [...]uint8{0, 4, 10, 15, 23, 34, 41}

func (i Attribute) String() string {
	i -= 1
	if i < 0 || i >= Attribute(len(_Attribute_index)-1) {
		return "Attribute(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Attribute_name[_Attribute_index[i]:_Attribute_index[i+1]]
}
