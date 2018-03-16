// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type SignatureValueType struct {
	w3c.Base64Binary
	Id *w3c.ID `xml:"Id,attr,omitempty"`
}

func NewSignatureValueType() *w3c.Base64Binary {
	return new(w3c.Base64Binary)
}

func (me *SignatureValueType) SetId(value *w3c.ID) {
	me.Id = value
}
