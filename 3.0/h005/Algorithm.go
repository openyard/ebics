// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
