// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type CanonicalizationMethodType struct {
	Algorithm *w3c.AnyURI `xml:"Algorithm,attr"`

	Any []*w3c.Any
}

func NewCanonicalizationMethodType() *CanonicalizationMethodType {
	return new(CanonicalizationMethodType)
}

func (me *CanonicalizationMethodType) SetAlgorithm(value *w3c.AnyURI) {
	me.Algorithm = value
}
