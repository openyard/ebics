// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h000

// ComplexElement
type HEVResponseDataTypeVersionNumber struct {
	Value           *VersionNumberType   `xml:",chardata"`
	ProtocolVersion *ProtocolVersionType `xml:"ProtocolVersion,attr"`
}

func NewHEVResponseDataTypeVersionNumber() *HEVResponseDataTypeVersionNumber {
	return new(HEVResponseDataTypeVersionNumber)
}

func (me *HEVResponseDataTypeVersionNumber) SetProtocolVersion(value *ProtocolVersionType) {
	me.ProtocolVersion = value
}
