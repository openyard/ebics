// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type NumSegmentsType w3c.NonNegativeInteger

func NewNumSegmentsType(value *w3c.NonNegativeInteger) *NumSegmentsType {
	me := (*NumSegmentsType)(value)
	return me
}
