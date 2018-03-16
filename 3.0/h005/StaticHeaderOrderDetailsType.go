// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type StaticHeaderOrderDetailsType struct {
	AdminOrderType *StaticHeaderOrderDetailsTypeAdminOrderType `xml:"AdminOrderType"`
	OrderID        *OrderIDType                                `xml:"OrderID,omitempty"`
	OrderParams    *OrderParams                                `xml:"OrderParams"`
}

func NewStaticHeaderOrderDetailsType() *StaticHeaderOrderDetailsType {
	return new(StaticHeaderOrderDetailsType)
}

func (me *StaticHeaderOrderDetailsType) SetAdminOrderType(value *StaticHeaderOrderDetailsTypeAdminOrderType) {
	me.AdminOrderType = value
}

func (me *StaticHeaderOrderDetailsType) AddAdminOrderType() *StaticHeaderOrderDetailsTypeAdminOrderType {
	me.AdminOrderType = new(StaticHeaderOrderDetailsTypeAdminOrderType)
	return me.AdminOrderType
}

func (me *StaticHeaderOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *StaticHeaderOrderDetailsType) AddOrderID() *OrderIDType {
	me.OrderID = new(OrderIDType)
	return me.OrderID
}

func (me *StaticHeaderOrderDetailsType) SetOrderParams(value *OrderParams) {
	me.OrderParams = value
}
