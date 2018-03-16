// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

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
