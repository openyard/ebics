// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HCARequestOrderDataType struct {
	AuthenticationPubKeyInfo *AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     *EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	PartnerID                *PartnerIDType                `xml:"PartnerID"`
	UserID                   *UserIDType                   `xml:"UserID"`

	Any []*w3c.Any
}

func NewHCARequestOrderDataType() *HCARequestOrderDataType {
	return new(HCARequestOrderDataType)
}

func (me *HCARequestOrderDataType) SetAuthenticationPubKeyInfo(value *AuthenticationPubKeyInfoType) {
	me.AuthenticationPubKeyInfo = value
}

func (me *HCARequestOrderDataType) AddAuthenticationPubKeyInfo() *AuthenticationPubKeyInfoType {
	me.AuthenticationPubKeyInfo = new(AuthenticationPubKeyInfoType)
	return me.AuthenticationPubKeyInfo
}

func (me *HCARequestOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HCARequestOrderDataType) AddEncryptionPubKeyInfo() *EncryptionPubKeyInfoType {
	me.EncryptionPubKeyInfo = new(EncryptionPubKeyInfoType)
	return me.EncryptionPubKeyInfo
}

func (me *HCARequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HCARequestOrderDataType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *HCARequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *HCARequestOrderDataType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}
