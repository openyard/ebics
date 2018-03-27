// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// AttributeGroup
type VersionAttrGroup struct {
	Version *ProtocolVersionType `xml:"Version,attr"`

	Revision *ProtocolRevisionType `xml:"Revision,attr,omitempty"`
}

func NewVersionAttrGroup() *VersionAttrGroup {
	return new(VersionAttrGroup)
}

func (me *VersionAttrGroup) SetVersion(value *ProtocolVersionType) {
	me.Version = value
}

func (me *VersionAttrGroup) SetRevision(value *ProtocolRevisionType) {
	me.Revision = value
}
