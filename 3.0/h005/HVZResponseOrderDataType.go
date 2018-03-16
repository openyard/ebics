// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
