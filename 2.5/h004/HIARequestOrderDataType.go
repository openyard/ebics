// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HIARequestOrderDataType struct {
	AuthenticationPubKeyInfo *AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     *EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	PartnerID                *PartnerIDType                `xml:"PartnerID"`
	UserID                   *UserIDType                   `xml:"UserID"`

	Any []*w3c.Any
}

func NewHIARequestOrderDataType() *HIARequestOrderDataType {
	return new(HIARequestOrderDataType)
}

func (me *HIARequestOrderDataType) SetAuthenticationPubKeyInfo(value *AuthenticationPubKeyInfoType) {
	me.AuthenticationPubKeyInfo = value
}

func (me *HIARequestOrderDataType) AddAuthenticationPubKeyInfo() *AuthenticationPubKeyInfoType {
	me.AuthenticationPubKeyInfo = new(AuthenticationPubKeyInfoType)
	return me.AuthenticationPubKeyInfo
}

func (me *HIARequestOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HIARequestOrderDataType) AddEncryptionPubKeyInfo() *EncryptionPubKeyInfoType {
	me.EncryptionPubKeyInfo = new(EncryptionPubKeyInfoType)
	return me.EncryptionPubKeyInfo
}

func (me *HIARequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HIARequestOrderDataType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *HIARequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *HIARequestOrderDataType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}
