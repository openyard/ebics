// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *TransformType) AddXPath() *w3c.String {
	me.XPath = new(w3c.String)
	return me.XPath
}
