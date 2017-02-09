// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type AccountType struct {
	Currency              CurrencyBaseType                 `xml:"Currency,attr,omitempty"`
	Description           AccountDescriptionType           `xml:"Description,attr,omitempty"`
	AccountHolder         AccountHolderType                `xml:"AccountHolder,omitempty"`
	AccountNumber         AccountTypeAccountNumber         `xml:"AccountNumber"`
	NationalAccountNumber AccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
	BankCode              AccountTypeBankCode              `xml:"BankCode"`
	NationalBankCode      AccountTypeNationalBankCode      `xml:"NationalBankCode"`
}
