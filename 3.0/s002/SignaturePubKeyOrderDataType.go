// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"

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

func (me *SignaturePubKeyOrderDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *SignaturePubKeyOrderDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}
