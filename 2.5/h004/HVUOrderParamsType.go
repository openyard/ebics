// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVUOrderParamsType struct {
	OrderTypes *OrderTListType `xml:"OrderTypes,omitempty"`

	Any []*w3c.Any
}

func NewHVUOrderParamsType() *HVUOrderParamsType {
	return new(HVUOrderParamsType)
}

func (me *HVUOrderParamsType) SetOrderTypes(value *OrderTListType) {
	me.OrderTypes = value
}

func (me *HVUOrderParamsType) AddOrderTypes() *OrderTListType {
	me.OrderTypes = new(OrderTListType)
	return me.OrderTypes
}
