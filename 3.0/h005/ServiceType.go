// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

// complex type
type ServiceType struct {
	ServiceName   ServiceNameStringType   `xml:"ServiceName,omitempty"`
	Scope         ScopeStringType         `xml:"Scope,omitempty"`
	ServiceOption ServiceOptionStringType `xml:"ServiceOption,omitempty"`
	Container     ContainerFlagType       `xml:"Container,omitempty"`
	MsgName       MessageType             `xml:"MsgName,omitempty"`
}
