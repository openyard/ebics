// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type SignatureMethodType struct {
	Algorithm        *w3c.AnyURI           `xml:"Algorithm,attr"`
	HMACOutputLength *HMACOutputLengthType `xml:"HMACOutputLength,omitempty"`

	Any []*w3c.Any
}

func NewSignatureMethodType() *SignatureMethodType {
	return new(SignatureMethodType)
}

func (me *SignatureMethodType) SetAlgorithm(value *w3c.AnyURI) {
	me.Algorithm = value
}

func (me *SignatureMethodType) SetHMACOutputLength(value *HMACOutputLengthType) {
	me.HMACOutputLength = value
}

func (me *SignatureMethodType) AddHMACOutputLength() *HMACOutputLengthType {
	me.HMACOutputLength = new(HMACOutputLengthType)
	return me.HMACOutputLength
}
