// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *HTDReponseOrderDataType) AddPartnerInfo() *PartnerInfoType {
	me.PartnerInfo = new(PartnerInfoType)
	return me.PartnerInfo
}

func (me *HTDReponseOrderDataType) SetUserInfo(value *UserInfoType) {
	me.UserInfo = value
}

func (me *HTDReponseOrderDataType) AddUserInfo() *UserInfoType {
	me.UserInfo = new(UserInfoType)
	return me.UserInfo
}
