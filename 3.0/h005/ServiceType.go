// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

// complex type
type ServiceType struct {
	ServiceName   ServiceNameStringType   `xml:"ServiceName,omitempty"`
	Scope         ScopeStringType         `xml:"Scope,omitempty"`
	ServiceOption ServiceOptionStringType `xml:"ServiceOption,omitempty"`
	Container     ContainerFlagType       `xml:"Container,omitempty"`
	MsgName       MessageType             `xml:"MsgName,omitempty"`
}
