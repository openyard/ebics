// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type ReadyToBeSigned struct {
	Value *w3c.Boolean `xml:"readyToBeSigned,attr"`
}

func NewReadyToBeSigned() *ReadyToBeSigned {
	return new(ReadyToBeSigned)
}

func (me *ReadyToBeSigned) SetValue(value *w3c.Boolean) {
	me.Value = value
}
