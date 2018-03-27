// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import esig "github.com/openyard/ebics/3.0/s002"
import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HCSRequestOrderDataType struct {
	AuthenticationPubKeyInfo *AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     *EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	PartnerID                *PartnerIDType                `xml:"PartnerID"`
	UserID                   *UserIDType                   `xml:"UserID"`
	SignaturePubKeyInfo      *esig.SignaturePubKeyInfo     `xml:"http://www.ebics.org/S002 SignaturePubKeyInfo"`

	Any []*w3c.Any
}

func NewHCSRequestOrderDataType() *HCSRequestOrderDataType {
	return new(HCSRequestOrderDataType)
}

func (me *HCSRequestOrderDataType) SetAuthenticationPubKeyInfo(value *AuthenticationPubKeyInfoType) {
	me.AuthenticationPubKeyInfo = value
}

func (me *HCSRequestOrderDataType) AddAuthenticationPubKeyInfo() *AuthenticationPubKeyInfoType {
	me.AuthenticationPubKeyInfo = new(AuthenticationPubKeyInfoType)
	return me.AuthenticationPubKeyInfo
}

func (me *HCSRequestOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HCSRequestOrderDataType) AddEncryptionPubKeyInfo() *EncryptionPubKeyInfoType {
	me.EncryptionPubKeyInfo = new(EncryptionPubKeyInfoType)
	return me.EncryptionPubKeyInfo
}

func (me *HCSRequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HCSRequestOrderDataType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *HCSRequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *HCSRequestOrderDataType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}

func (me *HCSRequestOrderDataType) SetSignaturePubKeyInfo(value *esig.SignaturePubKeyInfo) {
	me.SignaturePubKeyInfo = value
}
