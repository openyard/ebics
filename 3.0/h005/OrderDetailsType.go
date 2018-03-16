// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type OrderDetailsType struct {
	AdminOrderType *OrderTBaseType `xml:"AdminOrderType"`
}

func NewOrderDetailsType() *OrderDetailsType {
	return new(OrderDetailsType)
}

func (me *OrderDetailsType) SetAdminOrderType(value *OrderTBaseType) {
	me.AdminOrderType = value
}
