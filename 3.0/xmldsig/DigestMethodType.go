// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
