// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type UsageOrderType struct {
	Service []*RestrictedServiceType `xml:"Service,omitempty"`
}

func NewUsageOrderType() *UsageOrderType {
	return new(UsageOrderType)
}

func (me *UsageOrderType) AddService(value *RestrictedServiceType) {
	me.Service = append(me.Service, value)
}
