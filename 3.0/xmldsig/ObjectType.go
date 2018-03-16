// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
