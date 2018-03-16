// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *StaticHeaderOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *StaticHeaderOrderDetailsType) SetOrderParams(value *OrderParams) {
	me.OrderParams = value
}
