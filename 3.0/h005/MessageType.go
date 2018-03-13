// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

// complex type
type MessageType struct {
	MessageNameStringType
	Variant NumStringType     `xml:"variant,attr,omitempty"`
	Version NumStringType     `xml:"version,attr,omitempty"`
	Format  MessageTypeFormat `xml:"format,attr"`
}
