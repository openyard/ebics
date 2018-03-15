// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

// complex type
type MessageType struct {
	MessageNameStringType
	Variant *NumStringType     `xml:"variant,attr,omitempty"`
	Version *NumStringType     `xml:"version,attr,omitempty"`
	Format  *MessageTypeFormat `xml:"format,attr"`
}
