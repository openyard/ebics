// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
