// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type X509DataType struct {
	X509IssuerSerial X509IssuerSerialType `xml:"X509IssuerSerial"`
	X509SKI          w3c.Base64Binary     `xml:"X509SKI"`
	X509SubjectName  w3c.String           `xml:"X509SubjectName"`
	X509Certificate  w3c.Base64Binary     `xml:"X509Certificate"`
	X509CRL          w3c.Base64Binary     `xml:"X509CRL"`
}
