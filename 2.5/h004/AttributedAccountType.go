// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

// complex type
type AttributedAccountType struct {
	Currency              *CurrencyBaseType                           `xml:"Currency,attr,omitempty"`
	Description           *AccountDescriptionType                     `xml:"Description,attr"`
	AccountHolder         *AttributedAccountTypeAccountHolder         `xml:"AccountHolder,omitempty"`
	AccountNumber         *AttributedAccountTypeAccountNumber         `xml:"AccountNumber"`
	NationalAccountNumber *AttributedAccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
	BankCode              *AttributedAccountTypeBankCode              `xml:"BankCode"`
	NationalBankCode      *AttributedAccountTypeNationalBankCode      `xml:"NationalBankCode"`
}
