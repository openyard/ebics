// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type X509SKI w3c.Base64Binary

func NewX509SKI(value *w3c.Base64Binary) *X509SKI {
	me := (*X509SKI)(value)
	return me
}
