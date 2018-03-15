// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type X509IssuerSerialType struct {
	X509IssuerName   *w3c.String  `xml:"X509IssuerName"`
	X509SerialNumber *w3c.Integer `xml:"X509SerialNumber"`
}
