// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
