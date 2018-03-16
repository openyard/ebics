// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

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
