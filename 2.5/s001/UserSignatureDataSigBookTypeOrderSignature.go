// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
