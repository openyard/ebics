// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HAAResponseOrderDataType struct {
	Service []*RestrictedServiceType `xml:"Service,omitempty"`

	Any []*w3c.Any
}

func NewHAAResponseOrderDataType() *HAAResponseOrderDataType {
	return new(HAAResponseOrderDataType)
}

func (me *HAAResponseOrderDataType) AddService(value *RestrictedServiceType) {
	me.Service = append(me.Service, value)
}
