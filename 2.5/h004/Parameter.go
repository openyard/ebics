// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type Parameter struct {
	Name  *w3c.Token      `xml:"Name"`
	Value *ParameterValue `xml:"Value"`
}

func NewParameter() *Parameter {
	return new(Parameter)
}

func (me *Parameter) SetName(value *w3c.Token) {
	me.Name = value
}

func (me *Parameter) AddName() *w3c.Token {
	me.Name = new(w3c.Token)
	return me.Name
}

func (me *Parameter) SetValue(value *ParameterValue) {
	me.Value = value
}

func (me *Parameter) AddValue() *ParameterValue {
	me.Value = new(ParameterValue)
	return me.Value
}
