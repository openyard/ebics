// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type X509DataType struct {
	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial"`
	X509SKI          *w3c.Base64Binary     `xml:"X509SKI"`
	X509SubjectName  *w3c.String           `xml:"X509SubjectName"`
	X509Certificate  *w3c.Base64Binary     `xml:"X509Certificate"`
	X509CRL          *w3c.Base64Binary     `xml:"X509CRL"`
}
