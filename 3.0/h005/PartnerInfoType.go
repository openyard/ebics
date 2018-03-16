// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *PartnerInfoType) SetBankInfo(value *BankInfoType) {
	me.BankInfo = value
}

func (me *PartnerInfoType) AddAccountInfo(value *PartnerInfoTypeAccountInfo) {
	me.AccountInfo = append(me.AccountInfo, value)
}

func (me *PartnerInfoType) AddOrderInfo(value *AuthOrderInfoType) {
	me.OrderInfo = append(me.OrderInfo, value)
}
