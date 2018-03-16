// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVZResponseOrderDataType struct {
	OrderDetails []*HVZOrderDetailsType `xml:"OrderDetails"`

	Any []*w3c.Any
}

func NewHVZResponseOrderDataType() *HVZResponseOrderDataType {
	return new(HVZResponseOrderDataType)
}

func (me *HVZResponseOrderDataType) AddOrderDetails(value *HVZOrderDetailsType) {
	me.OrderDetails = append(me.OrderDetails, value)
}
