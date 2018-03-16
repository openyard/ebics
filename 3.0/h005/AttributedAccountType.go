// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type AttributedAccountType struct {
	Currency              *CurrencyBaseType                           `xml:"Currency,attr,omitempty"`
	Description           *AccountDescriptionType                     `xml:"Description,attr"`
	AccountHolder         *AttributedAccountTypeAccountHolder         `xml:"AccountHolder,omitempty"`
	AccountNumber         *AttributedAccountTypeAccountNumber         `xml:"AccountNumber"`
	NationalAccountNumber *AttributedAccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
	BankCode              *AttributedAccountTypeBankCode              `xml:"BankCode"`
	NationalBankCode      *AttributedAccountTypeNationalBankCode      `xml:"NationalBankCode"`
}

func NewAttributedAccountType() *AttributedAccountType {
	return new(AttributedAccountType)
}

func (me *AttributedAccountType) SetCurrency(value *CurrencyBaseType) {
	me.Currency = value
}

func (me *AttributedAccountType) SetDescription(value *AccountDescriptionType) {
	me.Description = value
}

func (me *AttributedAccountType) SetAccountHolder(value *AttributedAccountTypeAccountHolder) {
	me.AccountHolder = value
}

func (me *AttributedAccountType) SetAccountNumber(value *AttributedAccountTypeAccountNumber) {
	me.AccountNumber = value
}

func (me *AttributedAccountType) SetNationalAccountNumber(value *AttributedAccountTypeNationalAccountNumber) {
	me.NationalAccountNumber = value
}

func (me *AttributedAccountType) SetBankCode(value *AttributedAccountTypeBankCode) {
	me.BankCode = value
}

func (me *AttributedAccountType) SetNationalBankCode(value *AttributedAccountTypeNationalBankCode) {
	me.NationalBankCode = value
}
