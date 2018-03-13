// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

// complex type
type AccountType struct {
    
    Currency CurrencyBaseType `xml:"Currency,attr,omitempty"`
    Description AccountDescriptionType `xml:"Description,attr,omitempty"`
    AccountHolder AccountHolderType `xml:"AccountHolder,omitempty"`
    AccountNumber AccountTypeAccountNumber `xml:"AccountNumber"`
    NationalAccountNumber AccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
    BankCode AccountTypeBankCode `xml:"BankCode"`
    NationalBankCode AccountTypeNationalBankCode `xml:"NationalBankCode"`
    }
