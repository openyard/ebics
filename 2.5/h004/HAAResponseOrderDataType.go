// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HAAResponseOrderDataType struct {
	OrderTypes *OrderTListType `xml:"OrderTypes"`

	Any []*w3c.Any
}

func NewHAAResponseOrderDataType() *HAAResponseOrderDataType {
	return new(HAAResponseOrderDataType)
}

func (me *HAAResponseOrderDataType) SetOrderTypes(value *OrderTListType) {
	me.OrderTypes = value
}

func (me *HAAResponseOrderDataType) AddOrderTypes() *OrderTListType {
	me.OrderTypes = new(OrderTListType)
	return me.OrderTypes
}
