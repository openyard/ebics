// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type StaticHeaderOrderDetailsType struct {
	OrderType      *StaticHeaderOrderDetailsTypeOrderType `xml:"OrderType"`
	OrderID        *OrderIDType                           `xml:"OrderID,omitempty"`
	OrderAttribute *OrderAttributeType                    `xml:"OrderAttribute"`
	OrderParams    *OrderParams                           `xml:"OrderParams"`
}

func NewStaticHeaderOrderDetailsType() *StaticHeaderOrderDetailsType {
	return new(StaticHeaderOrderDetailsType)
}

func (me *StaticHeaderOrderDetailsType) SetOrderType(value *StaticHeaderOrderDetailsTypeOrderType) {
	me.OrderType = value
}

func (me *StaticHeaderOrderDetailsType) AddOrderType() *StaticHeaderOrderDetailsTypeOrderType {
	me.OrderType = new(StaticHeaderOrderDetailsTypeOrderType)
	return me.OrderType
}

func (me *StaticHeaderOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *StaticHeaderOrderDetailsType) AddOrderID() *OrderIDType {
	me.OrderID = new(OrderIDType)
	return me.OrderID
}

func (me *StaticHeaderOrderDetailsType) SetOrderAttribute(value *OrderAttributeType) {
	me.OrderAttribute = value
}

func (me *StaticHeaderOrderDetailsType) AddOrderAttribute() *OrderAttributeType {
	me.OrderAttribute = new(OrderAttributeType)
	return me.OrderAttribute
}

func (me *StaticHeaderOrderDetailsType) SetOrderParams(value *OrderParams) {
	me.OrderParams = value
}
