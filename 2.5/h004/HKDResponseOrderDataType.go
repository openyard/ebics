// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HKDResponseOrderDataType struct {
	PartnerInfo *PartnerInfoType `xml:"PartnerInfo"`
	UserInfo    []*UserInfoType  `xml:"UserInfo"`

	Any []*w3c.Any
}

func NewHKDResponseOrderDataType() *HKDResponseOrderDataType {
	return new(HKDResponseOrderDataType)
}

func (me *HKDResponseOrderDataType) SetPartnerInfo(value *PartnerInfoType) {
	me.PartnerInfo = value
}

func (me *HKDResponseOrderDataType) AddPartnerInfo() *PartnerInfoType {
	me.PartnerInfo = new(PartnerInfoType)
	return me.PartnerInfo
}

func (me *HKDResponseOrderDataType) AddUserInfo(value *UserInfoType) {
	me.UserInfo = append(me.UserInfo, value)
}
