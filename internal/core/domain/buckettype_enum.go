// Code generated by "go-enum -type=BucketType -linecomment"; DO NOT EDIT.

// Install go-enum by `go get install github.com/searKing/golang/tools/go-enum`
package domain

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BucketTypeInvalid-0]
	_ = x[BucketTypeAWS-1]
	_ = x[BucketTypeAzure-2]
	_ = x[BucketTypeOCI-3]
}

const _BucketType_name = "invalidawsazureoci"

var _BucketType_index = [...]uint8{0, 7, 10, 15, 18}

func _() {
	var _nil_BucketType_value = func() (val BucketType) { return }()

	// An "cannot convert BucketType literal (type BucketType) to type fmt.Stringer" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ fmt.Stringer = _nil_BucketType_value
}

func (i BucketType) String() string {
	if i < 0 || i >= BucketType(len(_BucketType_index)-1) {
		return "BucketType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BucketType_name[_BucketType_index[i]:_BucketType_index[i+1]]
}

// New returns a pointer to a new addr filled with the BucketType value passed in.
func (i BucketType) New() *BucketType {
	clone := i
	return &clone
}

var _BucketType_values = []BucketType{0, 1, 2, 3}

var _BucketType_name_to_values = map[string]BucketType{
	_BucketType_name[0:7]:   0,
	_BucketType_name[7:10]:  1,
	_BucketType_name[10:15]: 2,
	_BucketType_name[15:18]: 3,
}

// ParseBucketTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ParseBucketTypeString(s string) (BucketType, error) {
	if val, ok := _BucketType_name_to_values[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to BucketType values", s)
}

// BucketTypeValues returns all values of the enum
func BucketTypeValues() []BucketType {
	return _BucketType_values
}

// IsABucketType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i BucketType) Registered() bool {
	for _, v := range _BucketType_values {
		if i == v {
			return true
		}
	}
	return false
}

func _() {
	var _nil_BucketType_value = func() (val BucketType) { return }()

	// An "cannot convert BucketType literal (type BucketType) to type encoding.BinaryMarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.BinaryMarshaler = &_nil_BucketType_value

	// An "cannot convert BucketType literal (type BucketType) to type encoding.BinaryUnmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.BinaryUnmarshaler = &_nil_BucketType_value
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for BucketType
func (i BucketType) MarshalBinary() (data []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for BucketType
func (i *BucketType) UnmarshalBinary(data []byte) error {
	var err error
	*i, err = ParseBucketTypeString(string(data))
	return err
}

func _() {
	var _nil_BucketType_value = func() (val BucketType) { return }()

	// An "cannot convert BucketType literal (type BucketType) to type json.Marshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ json.Marshaler = _nil_BucketType_value

	// An "cannot convert BucketType literal (type BucketType) to type encoding.Unmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ json.Unmarshaler = &_nil_BucketType_value
}

// MarshalJSON implements the json.Marshaler interface for BucketType
func (i BucketType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for BucketType
func (i *BucketType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BucketType should be a string, got %s", data)
	}

	var err error
	*i, err = ParseBucketTypeString(s)
	return err
}

func _() {
	var _nil_BucketType_value = func() (val BucketType) { return }()

	// An "cannot convert BucketType literal (type BucketType) to type encoding.TextMarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.TextMarshaler = _nil_BucketType_value

	// An "cannot convert BucketType literal (type BucketType) to type encoding.TextUnmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.TextUnmarshaler = &_nil_BucketType_value
}

// MarshalText implements the encoding.TextMarshaler interface for BucketType
func (i BucketType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for BucketType
func (i *BucketType) UnmarshalText(text []byte) error {
	var err error
	*i, err = ParseBucketTypeString(string(text))
	return err
}

//func _() {
//	var _nil_BucketType_value = func() (val BucketType) { return }()
//
//	// An "cannot convert BucketType literal (type BucketType) to type yaml.Marshaler" compiler error signifies that the base type have changed.
//	// Re-run the go-enum command to generate them again.
//	var _ yaml.Marshaler = _nil_BucketType_value
//
//	// An "cannot convert BucketType literal (type BucketType) to type yaml.Unmarshaler" compiler error signifies that the base type have changed.
//	// Re-run the go-enum command to generate them again.
//	var _ yaml.Unmarshaler = &_nil_BucketType_value
//}

// MarshalYAML implements a YAML Marshaler for BucketType
func (i BucketType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for BucketType
func (i *BucketType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = ParseBucketTypeString(s)
	return err
}

func _() {
	var _nil_BucketType_value = func() (val BucketType) { return }()

	// An "cannot convert BucketType literal (type BucketType) to type driver.Valuer" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ driver.Valuer = _nil_BucketType_value

	// An "cannot convert BucketType literal (type BucketType) to type sql.Scanner" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ sql.Scanner = &_nil_BucketType_value
}

func (i BucketType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *BucketType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := ParseBucketTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

// BucketTypeSliceContains reports whether sunEnums is within enums.
func BucketTypeSliceContains(enums []BucketType, sunEnums ...BucketType) bool {
	var seenEnums = map[BucketType]bool{}
	for _, e := range sunEnums {
		seenEnums[e] = false
	}

	for _, v := range enums {
		if _, has := seenEnums[v]; has {
			seenEnums[v] = true
		}
	}

	for _, seen := range seenEnums {
		if !seen {
			return false
		}
	}

	return true
}

// BucketTypeSliceContainsAny reports whether any sunEnum is within enums.
func BucketTypeSliceContainsAny(enums []BucketType, sunEnums ...BucketType) bool {
	var seenEnums = map[BucketType]struct{}{}
	for _, e := range sunEnums {
		seenEnums[e] = struct{}{}
	}

	for _, v := range enums {
		if _, has := seenEnums[v]; has {
			return true
		}
	}

	return false
}