// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
