// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
