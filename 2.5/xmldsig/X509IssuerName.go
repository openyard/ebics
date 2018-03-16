// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type X509IssuerName w3c.String

func NewX509IssuerName(value *w3c.String) *X509IssuerName {
	me := (*X509IssuerName)(value)
	return me
}
