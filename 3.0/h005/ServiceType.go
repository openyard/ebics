// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

// complex type
type ServiceType struct {
	ServiceName   ServiceNameStringType   `xml:"ServiceName,omitempty"`
	Scope         ScopeStringType         `xml:"Scope,omitempty"`
	ServiceOption ServiceOptionStringType `xml:"ServiceOption,omitempty"`
	Container     ContainerFlagType       `xml:"Container,omitempty"`
	MsgName       MessageType             `xml:"MsgName,omitempty"`
}
