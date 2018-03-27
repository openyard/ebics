// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
