// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *HPBResponseOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HPBResponseOrderDataType) SetHostID(value *HostIDType) {
	me.HostID = value
}

func (me *HPBResponseOrderDataType) SetSignaturePubKeyInfo(value *esig.SignaturePubKeyInfo) {
	me.SignaturePubKeyInfo = value
}
