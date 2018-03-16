// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *OrderDetailsType) AddOrderType() *OrderTBaseType {
	me.OrderType = new(OrderTBaseType)
	return me.OrderType
}

func (me *OrderDetailsType) SetOrderAttribute(value *OrderAttributeBaseType) {
	me.OrderAttribute = value
}

func (me *OrderDetailsType) AddOrderAttribute() *OrderAttributeBaseType {
	me.OrderAttribute = new(OrderAttributeBaseType)
	return me.OrderAttribute
}
