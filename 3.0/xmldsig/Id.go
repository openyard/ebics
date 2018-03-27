// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// Attribute
type Id struct {
	Value *w3c.ID `xml:"Id,attr"`
}

func NewId() *Id {
	return new(Id)
}

func (me *Id) SetValue(value *w3c.ID) {
	me.Value = value
}
