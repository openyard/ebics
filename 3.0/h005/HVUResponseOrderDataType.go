// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVUResponseOrderDataType struct {
	OrderDetails []*HVUOrderDetailsType `xml:"OrderDetails"`

	Any []*w3c.Any
}

func NewHVUResponseOrderDataType() *HVUResponseOrderDataType {
	return new(HVUResponseOrderDataType)
}

func (me *HVUResponseOrderDataType) AddOrderDetails(value *HVUOrderDetailsType) {
	me.OrderDetails = append(me.OrderDetails, value)
}
