// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexElement
type HVTOrderInfoTypeAmount struct {
	Value    *AmountValueType  `xml:",chardata"`
	IsCredit *w3c.Boolean      `xml:"IsCredit,attr,omitempty"`
	Currency *CurrencyBaseType `xml:"Currency,attr,omitempty"`
}

func NewHVTOrderInfoTypeAmount() *HVTOrderInfoTypeAmount {
	return new(HVTOrderInfoTypeAmount)
}

func (me *HVTOrderInfoTypeAmount) SetIsCredit(value *w3c.Boolean) {
	me.IsCredit = value
}

func (me *HVTOrderInfoTypeAmount) SetCurrency(value *CurrencyBaseType) {
	me.Currency = value
}
