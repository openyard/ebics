// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type Type struct {
	Value *w3c.NCName `xml:"Type,attr"`
}

func NewType() *Type {
	return new(Type)
}

func (me *Type) SetValue(value *w3c.NCName) {
	me.Value = value
}
