// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
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
