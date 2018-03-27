// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// Attribute
type Currency struct {
	Value *CurrencyBaseType `xml:"Currency,attr"`
}

func NewCurrency() *Currency {
	return new(Currency)
}

func (me *Currency) SetValue(value *CurrencyBaseType) {
	me.Value = value
}
