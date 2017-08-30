// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
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
