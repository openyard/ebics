// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HVUOrderParamsType struct {
	ServiceFilter []*ServiceType `xml:"ServiceFilter,omitempty"`

	Any []*w3c.Any
}

func NewHVUOrderParamsType() *HVUOrderParamsType {
	return new(HVUOrderParamsType)
}

func (me *HVUOrderParamsType) AddServiceFilter(value *ServiceType) {
	me.ServiceFilter = append(me.ServiceFilter, value)
}
