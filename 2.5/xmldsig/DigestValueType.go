// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type DigestValueType w3c.Base64Binary

func NewDigestValueType(value *w3c.Base64Binary) *DigestValueType {
	me := (*DigestValueType)(value)
	return me
}
