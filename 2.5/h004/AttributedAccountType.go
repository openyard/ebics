// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type AttributedAccountType struct {
	Currency              CurrencyBaseType                           `xml:"Currency,attr,omitempty"`
	Description           AccountDescriptionType                     `xml:"Description,attr"`
	AccountHolder         AttributedAccountTypeAccountHolder         `xml:"AccountHolder,omitempty"`
	AccountNumber         AttributedAccountTypeAccountNumber         `xml:"AccountNumber"`
	NationalAccountNumber AttributedAccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
	BankCode              AttributedAccountTypeBankCode              `xml:"BankCode"`
	NationalBankCode      AttributedAccountTypeNationalBankCode      `xml:"NationalBankCode"`
}
