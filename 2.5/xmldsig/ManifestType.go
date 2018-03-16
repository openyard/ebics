// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
