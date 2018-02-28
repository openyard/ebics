// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type X509DataType struct {
	X509IssuerSerial X509IssuerSerialType `xml:"X509IssuerSerial"`
	X509SKI          w3c.Base64Binary     `xml:"X509SKI"`
	X509SubjectName  w3c.String           `xml:"X509SubjectName"`
	X509Certificate  w3c.Base64Binary     `xml:"X509Certificate"`
	X509CRL          w3c.Base64Binary     `xml:"X509CRL"`
}
