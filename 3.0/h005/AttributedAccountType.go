// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *AttributedAccountType) AddAccountHolder() *AttributedAccountTypeAccountHolder {
	me.AccountHolder = new(AttributedAccountTypeAccountHolder)
	return me.AccountHolder
}

func (me *AttributedAccountType) SetAccountNumber(value *AttributedAccountTypeAccountNumber) {
	me.AccountNumber = value
}

func (me *AttributedAccountType) AddAccountNumber() *AttributedAccountTypeAccountNumber {
	me.AccountNumber = new(AttributedAccountTypeAccountNumber)
	return me.AccountNumber
}

func (me *AttributedAccountType) SetNationalAccountNumber(value *AttributedAccountTypeNationalAccountNumber) {
	me.NationalAccountNumber = value
}

func (me *AttributedAccountType) AddNationalAccountNumber() *AttributedAccountTypeNationalAccountNumber {
	me.NationalAccountNumber = new(AttributedAccountTypeNationalAccountNumber)
	return me.NationalAccountNumber
}

func (me *AttributedAccountType) SetBankCode(value *AttributedAccountTypeBankCode) {
	me.BankCode = value
}

func (me *AttributedAccountType) AddBankCode() *AttributedAccountTypeBankCode {
	me.BankCode = new(AttributedAccountTypeBankCode)
	return me.BankCode
}

func (me *AttributedAccountType) SetNationalBankCode(value *AttributedAccountTypeNationalBankCode) {
	me.NationalBankCode = value
}

func (me *AttributedAccountType) AddNationalBankCode() *AttributedAccountTypeNationalBankCode {
	me.NationalBankCode = new(AttributedAccountTypeNationalBankCode)
	return me.NationalBankCode
}
