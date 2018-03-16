// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type X509DataType struct {
	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial"`
	X509SKI          *w3c.Base64Binary     `xml:"X509SKI"`
	X509SubjectName  *w3c.String           `xml:"X509SubjectName"`
	X509Certificate  *w3c.Base64Binary     `xml:"X509Certificate"`
	X509CRL          *w3c.Base64Binary     `xml:"X509CRL"`
}

func NewX509DataType() *X509DataType {
	return new(X509DataType)
}

func (me *X509DataType) SetX509IssuerSerial(value *X509IssuerSerialType) {
	me.X509IssuerSerial = value
}

func (me *X509DataType) SetX509SKI(value *w3c.Base64Binary) {
	me.X509SKI = value
}

func (me *X509DataType) SetX509SubjectName(value *w3c.String) {
	me.X509SubjectName = value
}

func (me *X509DataType) SetX509Certificate(value *w3c.Base64Binary) {
	me.X509Certificate = value
}

func (me *X509DataType) SetX509CRL(value *w3c.Base64Binary) {
	me.X509CRL = value
}
