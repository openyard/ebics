// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
