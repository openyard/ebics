// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVZResponseOrderDataType struct {
	OrderDetails []*HVZOrderDetailsType `xml:"OrderDetails"`

	Any []*w3c.Any
}

func NewHVZResponseOrderDataType() *HVZResponseOrderDataType {
	return new(HVZResponseOrderDataType)
}

func (me *HVZResponseOrderDataType) AddOrderDetails(value *HVZOrderDetailsType) {
	me.OrderDetails = append(me.OrderDetails, value)
}
