// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// attribute group
type VersionAttrGroup struct {
	Version ProtocolVersionType `xml:"Version,attr"`

	Revision ProtocolRevisionType `xml:"Revision,attr,omitempty"`
}
