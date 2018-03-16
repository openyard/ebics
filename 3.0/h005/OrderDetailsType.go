// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
