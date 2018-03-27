// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h000

// Attribute
type ProtocolVersion struct {
	Value *ProtocolVersionType `xml:"ProtocolVersion,attr"`
}

func NewProtocolVersion() *ProtocolVersion {
	return new(ProtocolVersion)
}

func (me *ProtocolVersion) SetValue(value *ProtocolVersionType) {
	me.Value = value
}
