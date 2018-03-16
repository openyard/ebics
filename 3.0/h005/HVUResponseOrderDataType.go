// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
