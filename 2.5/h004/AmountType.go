// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type AmountType struct {
	AmountValueType
	Currency *CurrencyBaseType `xml:"Currency,attr,omitempty"`
}

func NewAmountType() *AmountValueType {
	return new(AmountValueType)
}

func (me *AmountType) SetCurrency(value *CurrencyBaseType) {
	me.Currency = value
}
