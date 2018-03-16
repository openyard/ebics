// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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
