// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
