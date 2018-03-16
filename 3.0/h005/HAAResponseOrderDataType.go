// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
