// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

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
