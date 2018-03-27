// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HSARequestOrderDataType struct {
	AuthenticationPubKeyInfo *AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     *EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	PartnerID                *PartnerIDType                `xml:"PartnerID"`
	UserID                   *UserIDType                   `xml:"UserID"`

	Any []*w3c.Any
}

func NewHSARequestOrderDataType() *HSARequestOrderDataType {
	return new(HSARequestOrderDataType)
}

func (me *HSARequestOrderDataType) SetAuthenticationPubKeyInfo(value *AuthenticationPubKeyInfoType) {
	me.AuthenticationPubKeyInfo = value
}

func (me *HSARequestOrderDataType) AddAuthenticationPubKeyInfo() *AuthenticationPubKeyInfoType {
	me.AuthenticationPubKeyInfo = new(AuthenticationPubKeyInfoType)
	return me.AuthenticationPubKeyInfo
}

func (me *HSARequestOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HSARequestOrderDataType) AddEncryptionPubKeyInfo() *EncryptionPubKeyInfoType {
	me.EncryptionPubKeyInfo = new(EncryptionPubKeyInfoType)
	return me.EncryptionPubKeyInfo
}

func (me *HSARequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HSARequestOrderDataType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *HSARequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *HSARequestOrderDataType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}
