// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type X509SKI w3c.Base64Binary

func NewX509SKI(value *w3c.Base64Binary) *X509SKI {
	me := (*X509SKI)(value)
	return me
}
