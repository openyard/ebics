// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import esig "github.com/openyard/ebics/3.0/s002"
import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HPBResponseOrderDataType struct {
	AuthenticationPubKeyInfo *AuthenticationPubKeyInfoType `xml:"AuthenticationPubKeyInfo"`
	EncryptionPubKeyInfo     *EncryptionPubKeyInfoType     `xml:"EncryptionPubKeyInfo"`
	HostID                   *HostIDType                   `xml:"HostID"`
	SignaturePubKeyInfo      *esig.SignaturePubKeyInfo     `xml:"http://www.ebics.org/S002 SignaturePubKeyInfo,omitempty"`

	Any []*w3c.Any
}

func NewHPBResponseOrderDataType() *HPBResponseOrderDataType {
	return new(HPBResponseOrderDataType)
}

func (me *HPBResponseOrderDataType) SetAuthenticationPubKeyInfo(value *AuthenticationPubKeyInfoType) {
	me.AuthenticationPubKeyInfo = value
}

func (me *HPBResponseOrderDataType) AddAuthenticationPubKeyInfo() *AuthenticationPubKeyInfoType {
	me.AuthenticationPubKeyInfo = new(AuthenticationPubKeyInfoType)
	return me.AuthenticationPubKeyInfo
}

func (me *HPBResponseOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HPBResponseOrderDataType) AddEncryptionPubKeyInfo() *EncryptionPubKeyInfoType {
	me.EncryptionPubKeyInfo = new(EncryptionPubKeyInfoType)
	return me.EncryptionPubKeyInfo
}

func (me *HPBResponseOrderDataType) SetHostID(value *HostIDType) {
	me.HostID = value
}

func (me *HPBResponseOrderDataType) AddHostID() *HostIDType {
	me.HostID = new(HostIDType)
	return me.HostID
}

func (me *HPBResponseOrderDataType) SetSignaturePubKeyInfo(value *esig.SignaturePubKeyInfo) {
	me.SignaturePubKeyInfo = value
}
