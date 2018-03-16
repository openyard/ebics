// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// Attribute
type International struct {
	Value *w3c.Boolean `xml:"international,attr"`
}

func NewInternational() *International {
	return new(International)
}

func (me *International) SetValue(value *w3c.Boolean) {
	me.Value = value
}
