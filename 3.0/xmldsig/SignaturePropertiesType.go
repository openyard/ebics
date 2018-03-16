// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type SignaturePropertiesType struct {
	Id                *w3c.ID            `xml:"Id,attr,omitempty"`
	SignatureProperty *SignatureProperty `xml:"SignatureProperty"`
}

func NewSignaturePropertiesType() *SignaturePropertiesType {
	return new(SignaturePropertiesType)
}

func (me *SignaturePropertiesType) SetId(value *w3c.ID) {
	me.Id = value
}

func (me *SignaturePropertiesType) SetSignatureProperty(value *SignatureProperty) {
	me.SignatureProperty = value
}
