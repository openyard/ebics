// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type Algorithm struct {
	Value *w3c.AnyURI `xml:"Algorithm,attr"`
}

func NewAlgorithm() *Algorithm {
	return new(Algorithm)
}

func (me *Algorithm) SetValue(value *w3c.AnyURI) {
	me.Value = value
}
