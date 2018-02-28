// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

// attribute group
type VersionAttrGroup struct {
	Version ProtocolVersionType `xml:"Version,attr"`

	Revision ProtocolRevisionType `xml:"Revision,attr,omitempty"`
}
