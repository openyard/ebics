// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *X509IssuerSerialType) AddX509IssuerName() *w3c.String {
	me.X509IssuerName = new(w3c.String)
	return me.X509IssuerName
}

func (me *X509IssuerSerialType) SetX509SerialNumber(value *w3c.Integer) {
	me.X509SerialNumber = value
}

func (me *X509IssuerSerialType) AddX509SerialNumber() *w3c.Integer {
	me.X509SerialNumber = new(w3c.Integer)
	return me.X509SerialNumber
}
