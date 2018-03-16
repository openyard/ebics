// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
