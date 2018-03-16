// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type NumSigDone struct {
	Value *w3c.NonNegativeInteger `xml:"NumSigDone,attr"`
}

func NewNumSigDone() *NumSigDone {
	return new(NumSigDone)
}

func (me *NumSigDone) SetValue(value *w3c.NonNegativeInteger) {
	me.Value = value
}
