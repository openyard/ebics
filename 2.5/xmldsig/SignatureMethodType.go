// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
