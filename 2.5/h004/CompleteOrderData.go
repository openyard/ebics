// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type CompleteOrderData struct {
	Value *w3c.Boolean `xml:"completeOrderData,attr"`
}

func NewCompleteOrderData() *CompleteOrderData {
	return new(CompleteOrderData)
}

func (me *CompleteOrderData) SetValue(value *w3c.Boolean) {
	me.Value = value
}
