// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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

func (me *HCARequestOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HCARequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HCARequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}
