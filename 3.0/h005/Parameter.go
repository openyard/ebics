// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
