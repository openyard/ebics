// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type SignaturePropertyType struct {
	Target *w3c.AnyURI `xml:"Target,attr"`
	Id     *w3c.ID     `xml:"Id,attr,omitempty"`

	Any []*w3c.Any
}

func NewSignaturePropertyType() *SignaturePropertyType {
	return new(SignaturePropertyType)
}

func (me *SignaturePropertyType) SetTarget(value *w3c.AnyURI) {
	me.Target = value
}

func (me *SignaturePropertyType) SetId(value *w3c.ID) {
	me.Id = value
}
