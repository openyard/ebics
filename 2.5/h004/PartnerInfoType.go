// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type PartnerInfoType struct {
	AddressInfo AddressInfoType            `xml:"AddressInfo"`
	BankInfo    BankInfoType               `xml:"BankInfo"`
	AccountInfo PartnerInfoTypeAccountInfo `xml:"AccountInfo,omitempty"`
	OrderInfo   AuthOrderInfoType          `xml:"OrderInfo"`
}
