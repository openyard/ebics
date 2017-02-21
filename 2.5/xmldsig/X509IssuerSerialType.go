// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type X509IssuerSerialType struct {
	X509IssuerName   w3c.String  `xml:"X509IssuerName"`
	X509SerialNumber w3c.Integer `xml:"X509SerialNumber"`
}
