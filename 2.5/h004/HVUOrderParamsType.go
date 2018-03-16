// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
