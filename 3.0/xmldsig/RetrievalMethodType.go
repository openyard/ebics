// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type RetrievalMethodType struct {
	URI        *w3c.AnyURI `xml:"URI,attr"`
	Type       *w3c.AnyURI `xml:"Type,attr,omitempty"`
	Transforms *Transforms `xml:"Transforms,omitempty"`
}

func NewRetrievalMethodType() *RetrievalMethodType {
	return new(RetrievalMethodType)
}

func (me *RetrievalMethodType) SetURI(value *w3c.AnyURI) {
	me.URI = value
}

func (me *RetrievalMethodType) SetType(value *w3c.AnyURI) {
	me.Type = value
}

func (me *RetrievalMethodType) SetTransforms(value *Transforms) {
	me.Transforms = value
}
