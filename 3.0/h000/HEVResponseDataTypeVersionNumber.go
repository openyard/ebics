// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
