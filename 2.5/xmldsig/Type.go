// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type Type struct {
	Value *w3c.AnyURI `xml:"Type,attr"`
}

func NewType() *Type {
	return new(Type)
}

func (me *Type) SetValue(value *w3c.AnyURI) {
	me.Value = value
}
