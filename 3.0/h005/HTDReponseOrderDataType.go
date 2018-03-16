// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HTDReponseOrderDataType struct {
	PartnerInfo *PartnerInfoType `xml:"PartnerInfo"`
	UserInfo    *UserInfoType    `xml:"UserInfo"`

	Any []*w3c.Any
}

func NewHTDReponseOrderDataType() *HTDReponseOrderDataType {
	return new(HTDReponseOrderDataType)
}

func (me *HTDReponseOrderDataType) SetPartnerInfo(value *PartnerInfoType) {
	me.PartnerInfo = value
}

func (me *HTDReponseOrderDataType) SetUserInfo(value *UserInfoType) {
	me.UserInfo = value
}
