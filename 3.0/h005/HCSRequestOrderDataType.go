// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *HCSRequestOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HCSRequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HCSRequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *HCSRequestOrderDataType) SetSignaturePubKeyInfo(value *esig.SignaturePubKeyInfo) {
	me.SignaturePubKeyInfo = value
}
