// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
