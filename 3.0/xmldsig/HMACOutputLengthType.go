// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type HMACOutputLengthType w3c.Integer

func NewHMACOutputLengthType(value *w3c.Integer) *HMACOutputLengthType {
	me := (*HMACOutputLengthType)(value)
	return me
}
