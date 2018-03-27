// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
