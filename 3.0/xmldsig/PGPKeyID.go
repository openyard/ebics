// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type PGPKeyID w3c.Base64Binary

func NewPGPKeyID(value *w3c.Base64Binary) *PGPKeyID {
	me := (*PGPKeyID)(value)
	return me
}
