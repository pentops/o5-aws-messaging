// Code generated by protoc-gen-go-sugar. DO NOT EDIT.

package messaging_pb

import (
	driver "database/sql/driver"
	fmt "fmt"
)

type IsMessage_Extension = isMessage_Extension

// WireEncoding
const (
	WireEncoding_UNSPECIFIED WireEncoding = 0
	WireEncoding_PROTOJSON   WireEncoding = 1
	WireEncoding_RAW         WireEncoding = 2
	WireEncoding_J5_JSON     WireEncoding = 3
)

var (
	WireEncoding_name_short = map[int32]string{
		0: "UNSPECIFIED",
		1: "PROTOJSON",
		2: "RAW",
		3: "J5_JSON",
	}
	WireEncoding_value_short = map[string]int32{
		"UNSPECIFIED": 0,
		"PROTOJSON":   1,
		"RAW":         2,
		"J5_JSON":     3,
	}
	WireEncoding_value_either = map[string]int32{
		"UNSPECIFIED":               0,
		"WIRE_ENCODING_UNSPECIFIED": 0,
		"PROTOJSON":                 1,
		"WIRE_ENCODING_PROTOJSON":   1,
		"RAW":                       2,
		"WIRE_ENCODING_RAW":         2,
		"J5_JSON":                   3,
		"WIRE_ENCODING_J5_JSON":     3,
	}
)

// ShortString returns the un-prefixed string representation of the enum value
func (x WireEncoding) ShortString() string {
	return WireEncoding_name_short[int32(x)]
}
func (x WireEncoding) Value() (driver.Value, error) {
	return []uint8(x.ShortString()), nil
}
func (x *WireEncoding) Scan(value interface{}) error {
	var strVal string
	switch vt := value.(type) {
	case []uint8:
		strVal = string(vt)
	case string:
		strVal = vt
	default:
		return fmt.Errorf("invalid type %T", value)
	}
	val := WireEncoding_value_either[strVal]
	*x = WireEncoding(val)
	return nil
}
