// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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

func (me *HVTOrderParamsType) SetParameter(value *Parameter) {
	me.Parameter = value
}
