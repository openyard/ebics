// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
