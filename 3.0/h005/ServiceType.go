// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

// complex type
type ServiceType struct {
	ServiceName   *ServiceNameStringType   `xml:"ServiceName,omitempty"`
	Scope         *ScopeStringType         `xml:"Scope,omitempty"`
	ServiceOption *ServiceOptionStringType `xml:"ServiceOption,omitempty"`
	Container     *ContainerFlagType       `xml:"Container,omitempty"`
	MsgName       *MessageType             `xml:"MsgName,omitempty"`
}
