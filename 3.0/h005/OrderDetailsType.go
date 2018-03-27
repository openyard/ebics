// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *OrderDetailsType) AddAdminOrderType() *OrderTBaseType {
	me.AdminOrderType = new(OrderTBaseType)
	return me.AdminOrderType
}
