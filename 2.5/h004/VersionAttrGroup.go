// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
