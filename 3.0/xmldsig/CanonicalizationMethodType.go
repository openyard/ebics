// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

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
