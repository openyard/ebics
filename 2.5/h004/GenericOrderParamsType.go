// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
