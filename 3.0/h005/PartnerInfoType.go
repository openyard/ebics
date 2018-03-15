// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

// complex type
type PartnerInfoType struct {
	AddressInfo *AddressInfoType              `xml:"AddressInfo"`
	BankInfo    *BankInfoType                 `xml:"BankInfo"`
	AccountInfo []*PartnerInfoTypeAccountInfo `xml:"AccountInfo,omitempty"`
	OrderInfo   []*AuthOrderInfoType          `xml:"OrderInfo"`
}
