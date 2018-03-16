// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
