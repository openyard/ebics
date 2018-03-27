// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVTOrderParamsType struct {
	OrderFlags *HVTOrderParamsTypeOrderFlags `xml:"OrderFlags"`
	Parameter  *Parameter                    `xml:"Parameter,omitempty"`

	Any []*w3c.Any
}

func NewHVTOrderParamsType() *HVTOrderParamsType {
	return new(HVTOrderParamsType)
}

func (me *HVTOrderParamsType) SetOrderFlags(value *HVTOrderParamsTypeOrderFlags) {
	me.OrderFlags = value
}

func (me *HVTOrderParamsType) AddOrderFlags() *HVTOrderParamsTypeOrderFlags {
	me.OrderFlags = new(HVTOrderParamsTypeOrderFlags)
	return me.OrderFlags
}

func (me *HVTOrderParamsType) SetParameter(value *Parameter) {
	me.Parameter = value
}
