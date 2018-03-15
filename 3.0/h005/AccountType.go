// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

// complex type
type AccountType struct {
	Currency              *CurrencyBaseType                 `xml:"Currency,attr,omitempty"`
	Description           *AccountDescriptionType           `xml:"Description,attr,omitempty"`
	AccountHolder         *AccountHolderType                `xml:"AccountHolder,omitempty"`
	AccountNumber         *AccountTypeAccountNumber         `xml:"AccountNumber"`
	NationalAccountNumber *AccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
	BankCode              *AccountTypeBankCode              `xml:"BankCode"`
	NationalBankCode      *AccountTypeNationalBankCode      `xml:"NationalBankCode"`
}
