// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
