// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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
