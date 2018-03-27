// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type SignatureValue w3c.Base64Binary

func NewSignatureValue(value *w3c.Base64Binary) *SignatureValue {
	me := (*SignatureValue)(value)
	return me
}
