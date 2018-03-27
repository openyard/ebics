// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type X509Certificate w3c.Base64Binary

func NewX509Certificate(value *w3c.Base64Binary) *X509Certificate {
	me := (*X509Certificate)(value)
	return me
}
