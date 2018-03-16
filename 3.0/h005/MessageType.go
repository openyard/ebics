// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type MessageType struct {
	MessageNameStringType
	Variant *NumStringType     `xml:"variant,attr,omitempty"`
	Version *NumStringType     `xml:"version,attr,omitempty"`
	Format  *MessageTypeFormat `xml:"format,attr"`
}

func NewMessageType() *MessageNameStringType {
	return new(MessageNameStringType)
}

func (me *MessageType) SetVariant(value *NumStringType) {
	me.Variant = value
}

func (me *MessageType) SetVersion(value *NumStringType) {
	me.Version = value
}

func (me *MessageType) SetFormat(value *MessageTypeFormat) {
	me.Format = value
}
