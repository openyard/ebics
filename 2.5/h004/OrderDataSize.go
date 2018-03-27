// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type OrderDataSize w3c.PositiveInteger

func NewOrderDataSize(value *w3c.PositiveInteger) *OrderDataSize {
	me := (*OrderDataSize)(value)
	return me
}
