// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
