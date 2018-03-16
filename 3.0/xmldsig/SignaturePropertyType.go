// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type SignaturePropertyType struct {
	Target *w3c.AnyURI `xml:"Target,attr"`
	Id     *w3c.ID     `xml:"Id,attr,omitempty"`

	Any []*w3c.Any
}

func NewSignaturePropertyType() *SignaturePropertyType {
	return new(SignaturePropertyType)
}

func (me *SignaturePropertyType) SetTarget(value *w3c.AnyURI) {
	me.Target = value
}

func (me *SignaturePropertyType) SetId(value *w3c.ID) {
	me.Id = value
}
