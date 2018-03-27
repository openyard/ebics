// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type PubKeyDigestType struct {
	DigestType
	Algorithm *w3c.AnyURI `xml:"Algorithm,attr"`
}

func NewPubKeyDigestType() *DigestType {
	return new(DigestType)
}

func (me *PubKeyDigestType) SetAlgorithm(value *w3c.AnyURI) {
	me.Algorithm = value
}
