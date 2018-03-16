// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *StaticHeaderOrderDetailsType) SetOrderID(value *OrderIDType) {
	me.OrderID = value
}

func (me *StaticHeaderOrderDetailsType) SetOrderAttribute(value *OrderAttributeType) {
	me.OrderAttribute = value
}

func (me *StaticHeaderOrderDetailsType) SetOrderParams(value *OrderParams) {
	me.OrderParams = value
}
