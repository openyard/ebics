// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import ds "github.com/openyard/ebics/3.0/xmldsig"
import w3c "github.com/openyard/ebics/3.0/w3c"

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
