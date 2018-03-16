// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type X509SerialNumber w3c.Integer

func NewX509SerialNumber(value *w3c.Integer) *X509SerialNumber {
	me := (*X509SerialNumber)(value)
	return me
}
