// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type OrderSignatureType w3c.Base64Binary

func NewOrderSignatureType(value *w3c.Base64Binary) *OrderSignatureType {
	me := (*OrderSignatureType)(value)
	return me
}
