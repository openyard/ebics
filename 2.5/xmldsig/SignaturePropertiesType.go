// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
