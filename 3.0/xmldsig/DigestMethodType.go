// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type DigestMethodType struct {
	Algorithm *w3c.AnyURI `xml:"Algorithm,attr"`

	Any []*w3c.Any
}

func NewDigestMethodType() *DigestMethodType {
	return new(DigestMethodType)
}

func (me *DigestMethodType) SetAlgorithm(value *w3c.AnyURI) {
	me.Algorithm = value
}
