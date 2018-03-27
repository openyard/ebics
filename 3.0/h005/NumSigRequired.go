// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
