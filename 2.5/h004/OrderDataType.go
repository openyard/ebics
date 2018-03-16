// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type OrderDataType w3c.Base64Binary

func NewOrderDataType(value *w3c.Base64Binary) *OrderDataType {
	me := (*OrderDataType)(value)
	return me
}
