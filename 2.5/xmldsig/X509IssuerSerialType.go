// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type X509IssuerSerialType struct {
	X509IssuerName   *w3c.String  `xml:"X509IssuerName"`
	X509SerialNumber *w3c.Integer `xml:"X509SerialNumber"`
}

func NewX509IssuerSerialType() *X509IssuerSerialType {
	return new(X509IssuerSerialType)
}

func (me *X509IssuerSerialType) SetX509IssuerName(value *w3c.String) {
	me.X509IssuerName = value
}

func (me *X509IssuerSerialType) SetX509SerialNumber(value *w3c.Integer) {
	me.X509SerialNumber = value
}
