// Generated with goxc v - rev
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type HVTOrderInfoTypeAmount struct {
	Value    AmountValueType  `xml:",chardata"`
	IsCredit w3c.Boolean      `xml:"IsCredit,attr,omitempty"`
	Currency CurrencyBaseType `xml:"Currency,attr,omitempty"`
}
