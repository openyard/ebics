// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type SignaturePubKeyOrderDataType struct {
	SignaturePubKeyInfo *SignaturePubKeyInfoType `xml:"SignaturePubKeyInfo"`
	PartnerID           *PartnerIDType           `xml:"PartnerID"`
	UserID              *UserIDType              `xml:"UserID"`

	Any []*w3c.Any
}

func NewSignaturePubKeyOrderDataType() *SignaturePubKeyOrderDataType {
	return new(SignaturePubKeyOrderDataType)
}

func (me *SignaturePubKeyOrderDataType) SetSignaturePubKeyInfo(value *SignaturePubKeyInfoType) {
	me.SignaturePubKeyInfo = value
}

func (me *SignaturePubKeyOrderDataType) AddSignaturePubKeyInfo() *SignaturePubKeyInfoType {
	me.SignaturePubKeyInfo = new(SignaturePubKeyInfoType)
	return me.SignaturePubKeyInfo
}

func (me *SignaturePubKeyOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *SignaturePubKeyOrderDataType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *SignaturePubKeyOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *SignaturePubKeyOrderDataType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}
