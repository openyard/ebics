// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type PartnerInfoType struct {
	AddressInfo *AddressInfoType              `xml:"AddressInfo"`
	BankInfo    *BankInfoType                 `xml:"BankInfo"`
	AccountInfo []*PartnerInfoTypeAccountInfo `xml:"AccountInfo,omitempty"`
	OrderInfo   []*AuthOrderInfoType          `xml:"OrderInfo"`
}

func NewPartnerInfoType() *PartnerInfoType {
	return new(PartnerInfoType)
}

func (me *PartnerInfoType) SetAddressInfo(value *AddressInfoType) {
	me.AddressInfo = value
}

func (me *PartnerInfoType) AddAddressInfo() *AddressInfoType {
	me.AddressInfo = new(AddressInfoType)
	return me.AddressInfo
}

func (me *PartnerInfoType) SetBankInfo(value *BankInfoType) {
	me.BankInfo = value
}

func (me *PartnerInfoType) AddBankInfo() *BankInfoType {
	me.BankInfo = new(BankInfoType)
	return me.BankInfo
}

func (me *PartnerInfoType) AddAccountInfo(value *PartnerInfoTypeAccountInfo) {
	me.AccountInfo = append(me.AccountInfo, value)
}

func (me *PartnerInfoType) AddOrderInfo(value *AuthOrderInfoType) {
	me.OrderInfo = append(me.OrderInfo, value)
}
