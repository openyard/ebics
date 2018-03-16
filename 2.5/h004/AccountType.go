// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type AccountType struct {
	Currency              *CurrencyBaseType                 `xml:"Currency,attr,omitempty"`
	Description           *AccountDescriptionType           `xml:"Description,attr,omitempty"`
	AccountHolder         *AccountHolderType                `xml:"AccountHolder,omitempty"`
	AccountNumber         *AccountTypeAccountNumber         `xml:"AccountNumber"`
	NationalAccountNumber *AccountTypeNationalAccountNumber `xml:"NationalAccountNumber"`
	BankCode              *AccountTypeBankCode              `xml:"BankCode"`
	NationalBankCode      *AccountTypeNationalBankCode      `xml:"NationalBankCode"`
}

func NewAccountType() *AccountType {
	return new(AccountType)
}

func (me *AccountType) SetCurrency(value *CurrencyBaseType) {
	me.Currency = value
}

func (me *AccountType) SetDescription(value *AccountDescriptionType) {
	me.Description = value
}

func (me *AccountType) SetAccountHolder(value *AccountHolderType) {
	me.AccountHolder = value
}

func (me *AccountType) SetAccountNumber(value *AccountTypeAccountNumber) {
	me.AccountNumber = value
}

func (me *AccountType) SetNationalAccountNumber(value *AccountTypeNationalAccountNumber) {
	me.NationalAccountNumber = value
}

func (me *AccountType) SetBankCode(value *AccountTypeBankCode) {
	me.BankCode = value
}

func (me *AccountType) SetNationalBankCode(value *AccountTypeNationalBankCode) {
	me.NationalBankCode = value
}
