// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

// complex type
type PartnerInfoType struct {
	AddressInfo AddressInfoType            `xml:"AddressInfo"`
	BankInfo    BankInfoType               `xml:"BankInfo"`
	AccountInfo PartnerInfoTypeAccountInfo `xml:"AccountInfo,omitempty"`
	OrderInfo   AuthOrderInfoType          `xml:"OrderInfo"`
}
