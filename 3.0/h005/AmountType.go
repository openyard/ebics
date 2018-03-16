// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
