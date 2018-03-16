// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type TransformType struct {
	Algorithm *w3c.AnyURI `xml:"Algorithm,attr"`
	XPath     *w3c.String `xml:"XPath"`

	Any []*w3c.Any
}

func NewTransformType() *TransformType {
	return new(TransformType)
}

func (me *TransformType) SetAlgorithm(value *w3c.AnyURI) {
	me.Algorithm = value
}

func (me *TransformType) SetXPath(value *w3c.String) {
	me.XPath = value
}
