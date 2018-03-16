// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type OrderDetailsType struct {
	OrderType      *OrderTBaseType         `xml:"OrderType"`
	OrderAttribute *OrderAttributeBaseType `xml:"OrderAttribute"`
}

func NewOrderDetailsType() *OrderDetailsType {
	return new(OrderDetailsType)
}

func (me *OrderDetailsType) SetOrderType(value *OrderTBaseType) {
	me.OrderType = value
}

func (me *OrderDetailsType) SetOrderAttribute(value *OrderAttributeBaseType) {
	me.OrderAttribute = value
}
