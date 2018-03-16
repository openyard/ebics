// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// Attribute
type NumSigRequired struct {
	Value *w3c.PositiveInteger `xml:"NumSigRequired,attr"`
}

func NewNumSigRequired() *NumSigRequired {
	return new(NumSigRequired)
}

func (me *NumSigRequired) SetValue(value *w3c.PositiveInteger) {
	me.Value = value
}
