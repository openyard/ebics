// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
