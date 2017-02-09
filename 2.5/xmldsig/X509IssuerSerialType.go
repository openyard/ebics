// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type X509IssuerSerialType struct {
	X509IssuerName   w3c.String  `xml:"X509IssuerName"`
	X509SerialNumber w3c.Integer `xml:"X509SerialNumber"`
}
