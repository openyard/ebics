// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVZOrderParamsType struct {
	OrderTypes *OrderTListType `xml:"OrderTypes,omitempty"`

	Any []*w3c.Any
}

func NewHVZOrderParamsType() *HVZOrderParamsType {
	return new(HVZOrderParamsType)
}

func (me *HVZOrderParamsType) SetOrderTypes(value *OrderTListType) {
	me.OrderTypes = value
}

func (me *HVZOrderParamsType) AddOrderTypes() *OrderTListType {
	me.OrderTypes = new(OrderTListType)
	return me.OrderTypes
}
