// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
