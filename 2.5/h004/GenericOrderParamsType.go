// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type GenericOrderParamsType struct {
	Parameter *Parameter `xml:"Parameter,omitempty"`
}

func NewGenericOrderParamsType() *GenericOrderParamsType {
	return new(GenericOrderParamsType)
}

func (me *GenericOrderParamsType) SetParameter(value *Parameter) {
	me.Parameter = value
}
