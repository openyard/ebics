// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type PreValidationAccountAuthType struct {
	AccountType
	Amount *AmountType `xml:"Amount,omitempty"`
}

func NewPreValidationAccountAuthType() *AccountType {
	return new(AccountType)
}

func (me *PreValidationAccountAuthType) SetAmount(value *AmountType) {
	me.Amount = value
}

func (me *PreValidationAccountAuthType) AddAmount() *AmountType {
	me.Amount = new(AmountType)
	return me.Amount
}
