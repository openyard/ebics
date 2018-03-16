// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type SignatureType struct {
	Id             *w3c.ID         `xml:"Id,attr,omitempty"`
	SignedInfo     *SignedInfo     `xml:"SignedInfo"`
	SignatureValue *SignatureValue `xml:"SignatureValue"`
	KeyInfo        *KeyInfo        `xml:"KeyInfo,omitempty"`
	Object         *Object         `xml:"Object,omitempty"`
}

func NewSignatureType() *SignatureType {
	return new(SignatureType)
}

func (me *SignatureType) SetId(value *w3c.ID) {
	me.Id = value
}

func (me *SignatureType) SetSignedInfo(value *SignedInfo) {
	me.SignedInfo = value
}

func (me *SignatureType) SetSignatureValue(value *SignatureValue) {
	me.SignatureValue = value
}

func (me *SignatureType) SetKeyInfo(value *KeyInfo) {
	me.KeyInfo = value
}

func (me *SignatureType) SetObject(value *Object) {
	me.Object = value
}
