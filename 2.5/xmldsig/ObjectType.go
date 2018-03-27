// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
