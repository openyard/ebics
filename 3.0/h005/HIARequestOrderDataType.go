// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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

func (me *HIARequestOrderDataType) SetEncryptionPubKeyInfo(value *EncryptionPubKeyInfoType) {
	me.EncryptionPubKeyInfo = value
}

func (me *HIARequestOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HIARequestOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}
