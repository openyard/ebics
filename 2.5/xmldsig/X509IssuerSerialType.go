// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type X509IssuerSerialType struct {
	X509IssuerName   w3c.String  `xml:"X509IssuerName"`
	X509SerialNumber w3c.Integer `xml:"X509SerialNumber"`
}
