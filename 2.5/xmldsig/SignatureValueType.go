// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type SignatureValueType struct {
	w3c.Base64Binary
	Id *w3c.ID `xml:"Id,attr,omitempty"`
}

func NewSignatureValueType() *w3c.Base64Binary {
	return new(w3c.Base64Binary)
}

func (me *SignatureValueType) SetId(value *w3c.ID) {
	me.Id = value
}
