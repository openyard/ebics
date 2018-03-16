// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

// ComplexType
type ServiceType struct {
	ServiceName   *ServiceNameStringType   `xml:"ServiceName,omitempty"`
	Scope         *ScopeStringType         `xml:"Scope,omitempty"`
	ServiceOption *ServiceOptionStringType `xml:"ServiceOption,omitempty"`
	Container     *ContainerFlagType       `xml:"Container,omitempty"`
	MsgName       *MessageType             `xml:"MsgName,omitempty"`
}

func NewServiceType() *ServiceType {
	return new(ServiceType)
}

func (me *ServiceType) SetServiceName(value *ServiceNameStringType) {
	me.ServiceName = value
}

func (me *ServiceType) SetScope(value *ScopeStringType) {
	me.Scope = value
}

func (me *ServiceType) SetServiceOption(value *ServiceOptionStringType) {
	me.ServiceOption = value
}

func (me *ServiceType) SetContainer(value *ContainerFlagType) {
	me.Container = value
}

func (me *ServiceType) SetMsgName(value *MessageType) {
	me.MsgName = value
}
