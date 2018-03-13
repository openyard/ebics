// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

// complex type
type AttributedAccountType struct {
    
    Currency CurrencyBaseType `xml:"Currency,attr,omitempty"`
    Description AccountDescriptionType `xml:"Description,attr"`
    AccountHolder AttributedAccountTypeAccountHolder `xml:"AccountHolder,omitempty"`
    AccountNumber AttributedAccountTypeAccountNumber `xml:"AccountNumber"`
    NationalAccountNumber AttributedAccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
    BankCode AttributedAccountTypeBankCode `xml:"BankCode"`
    NationalBankCode AttributedAccountTypeNationalBankCode `xml:"NationalBankCode"`
    }
