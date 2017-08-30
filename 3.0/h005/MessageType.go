// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

// complex type
type MessageType struct {
	MessageNameStringType
	Variant NumStringType     `xml:"variant,attr,omitempty"`
	Version NumStringType     `xml:"version,attr,omitempty"`
	Format  MessageTypeFormat `xml:"format,attr"`
}
