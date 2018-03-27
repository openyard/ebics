// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
