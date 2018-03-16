// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
