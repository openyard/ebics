// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
