// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type ObjectType struct {
	Id       *w3c.ID     `xml:"Id,attr,omitempty"`
	MimeType *w3c.String `xml:"MimeType,attr,omitempty"`
	Encoding *w3c.AnyURI `xml:"Encoding,attr,omitempty"`

	Any []*w3c.Any
}

func NewObjectType() *ObjectType {
	return new(ObjectType)
}

func (me *ObjectType) SetId(value *w3c.ID) {
	me.Id = value
}

func (me *ObjectType) SetMimeType(value *w3c.String) {
	me.MimeType = value
}

func (me *ObjectType) SetEncoding(value *w3c.AnyURI) {
	me.Encoding = value
}
