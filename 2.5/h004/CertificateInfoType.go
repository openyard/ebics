// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"
import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type CertificateInfoType struct {
	X509Data *ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data"`

	Any []*w3c.Any
}

func NewCertificateInfoType() *CertificateInfoType {
	return new(CertificateInfoType)
}

func (me *CertificateInfoType) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
