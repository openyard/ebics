// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type ManifestType struct {
	Id        *w3c.ID    `xml:"Id,attr,omitempty"`
	Reference *Reference `xml:"Reference"`
}

func NewManifestType() *ManifestType {
	return new(ManifestType)
}

func (me *ManifestType) SetId(value *w3c.ID) {
	me.Id = value
}

func (me *ManifestType) SetReference(value *Reference) {
	me.Reference = value
}
