// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

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
