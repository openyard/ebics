// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVZOrderParamsType struct {
	ServiceFilter []*ServiceType `xml:"ServiceFilter,omitempty"`

	Any []*w3c.Any
}

func NewHVZOrderParamsType() *HVZOrderParamsType {
	return new(HVZOrderParamsType)
}

func (me *HVZOrderParamsType) AddServiceFilter(value *ServiceType) {
	me.ServiceFilter = append(me.ServiceFilter, value)
}
