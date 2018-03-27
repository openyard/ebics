// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
