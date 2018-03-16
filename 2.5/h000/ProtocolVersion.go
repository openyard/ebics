// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
