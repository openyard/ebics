// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
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
