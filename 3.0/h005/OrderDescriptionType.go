// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type OrderDescriptionType w3c.NormalizedString

func NewOrderDescriptionType(value *w3c.NormalizedString) *OrderDescriptionType {
	me := (*OrderDescriptionType)(value)
	return me
}
