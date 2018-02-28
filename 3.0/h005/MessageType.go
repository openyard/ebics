// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

// complex type
type MessageType struct {
	MessageNameStringType
	Variant NumStringType     `xml:"variant,attr,omitempty"`
	Version NumStringType     `xml:"version,attr,omitempty"`
	Format  MessageTypeFormat `xml:"format,attr"`
}
