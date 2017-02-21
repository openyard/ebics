// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
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
