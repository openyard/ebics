// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type ReferenceType struct {
	Id           *w3c.ID       `xml:"Id,attr,omitempty"`
	URI          *w3c.AnyURI   `xml:"URI,attr,omitempty"`
	Type         *w3c.AnyURI   `xml:"Type,attr,omitempty"`
	Transforms   *Transforms   `xml:"Transforms,omitempty"`
	DigestMethod *DigestMethod `xml:"DigestMethod"`
	DigestValue  *DigestValue  `xml:"DigestValue"`
}

func NewReferenceType() *ReferenceType {
	return new(ReferenceType)
}

func (me *ReferenceType) SetId(value *w3c.ID) {
	me.Id = value
}

func (me *ReferenceType) SetURI(value *w3c.AnyURI) {
	me.URI = value
}

func (me *ReferenceType) SetType(value *w3c.AnyURI) {
	me.Type = value
}

func (me *ReferenceType) SetTransforms(value *Transforms) {
	me.Transforms = value
}

func (me *ReferenceType) SetDigestMethod(value *DigestMethod) {
	me.DigestMethod = value
}

func (me *ReferenceType) SetDigestValue(value *DigestValue) {
	me.DigestValue = value
}
