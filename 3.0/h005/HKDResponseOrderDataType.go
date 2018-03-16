// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
