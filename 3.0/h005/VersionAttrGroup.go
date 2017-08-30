// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

// attribute group
type VersionAttrGroup struct {
	Version ProtocolVersionType `xml:"Version,attr"`

	Revision ProtocolRevisionType `xml:"Revision,attr,omitempty"`
}
