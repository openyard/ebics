// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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

func (me *AccountType) AddAccountHolder() *AccountHolderType {
	me.AccountHolder = new(AccountHolderType)
	return me.AccountHolder
}

func (me *AccountType) SetAccountNumber(value *AccountTypeAccountNumber) {
	me.AccountNumber = value
}

func (me *AccountType) AddAccountNumber() *AccountTypeAccountNumber {
	me.AccountNumber = new(AccountTypeAccountNumber)
	return me.AccountNumber
}

func (me *AccountType) SetNationalAccountNumber(value *AccountTypeNationalAccountNumber) {
	me.NationalAccountNumber = value
}

func (me *AccountType) AddNationalAccountNumber() *AccountTypeNationalAccountNumber {
	me.NationalAccountNumber = new(AccountTypeNationalAccountNumber)
	return me.NationalAccountNumber
}

func (me *AccountType) SetBankCode(value *AccountTypeBankCode) {
	me.BankCode = value
}

func (me *AccountType) AddBankCode() *AccountTypeBankCode {
	me.BankCode = new(AccountTypeBankCode)
	return me.BankCode
}

func (me *AccountType) SetNationalBankCode(value *AccountTypeNationalBankCode) {
	me.NationalBankCode = value
}

func (me *AccountType) AddNationalBankCode() *AccountTypeNationalBankCode {
	me.NationalBankCode = new(AccountTypeNationalBankCode)
	return me.NationalBankCode
}
