// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type UserSignatureDataSigBookTypeOrderSignature struct {
	Value     *OrderSignatureType `xml:",chardata"`
	PartnerID *w3c.String         `xml:"PartnerID,attr"`
}

func NewUserSignatureDataSigBookTypeOrderSignature() *UserSignatureDataSigBookTypeOrderSignature {
	return new(UserSignatureDataSigBookTypeOrderSignature)
}

func (me *UserSignatureDataSigBookTypeOrderSignature) SetPartnerID(value *w3c.String) {
	me.PartnerID = value
}
