// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// Attribute
type Target struct {
	Value *w3c.AnyURI `xml:"Target,attr"`
}

func NewTarget() *Target {
	return new(Target)
}

func (me *Target) SetValue(value *w3c.AnyURI) {
	me.Value = value
}
