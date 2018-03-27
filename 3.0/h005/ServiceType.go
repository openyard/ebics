// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *ServiceType) AddServiceName() *ServiceNameStringType {
	me.ServiceName = new(ServiceNameStringType)
	return me.ServiceName
}

func (me *ServiceType) SetScope(value *ScopeStringType) {
	me.Scope = value
}

func (me *ServiceType) AddScope() *ScopeStringType {
	me.Scope = new(ScopeStringType)
	return me.Scope
}

func (me *ServiceType) SetServiceOption(value *ServiceOptionStringType) {
	me.ServiceOption = value
}

func (me *ServiceType) AddServiceOption() *ServiceOptionStringType {
	me.ServiceOption = new(ServiceOptionStringType)
	return me.ServiceOption
}

func (me *ServiceType) SetContainer(value *ContainerFlagType) {
	me.Container = value
}

func (me *ServiceType) AddContainer() *ContainerFlagType {
	me.Container = new(ContainerFlagType)
	return me.Container
}

func (me *ServiceType) SetMsgName(value *MessageType) {
	me.MsgName = value
}

func (me *ServiceType) AddMsgName() *MessageType {
	me.MsgName = new(MessageType)
	return me.MsgName
}
