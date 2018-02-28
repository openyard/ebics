// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
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
