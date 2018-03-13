// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type X509IssuerSerialType struct {
	X509IssuerName   w3c.String  `xml:"X509IssuerName"`
	X509SerialNumber w3c.Integer `xml:"X509SerialNumber"`
}
