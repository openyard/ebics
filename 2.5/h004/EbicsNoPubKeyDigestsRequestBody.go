// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"

// ComplexType
type EbicsNoPubKeyDigestsRequestBody struct {
	X509Data *ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}

func NewEbicsNoPubKeyDigestsRequestBody() *EbicsNoPubKeyDigestsRequestBody {
	return new(EbicsNoPubKeyDigestsRequestBody)
}

func (me *EbicsNoPubKeyDigestsRequestBody) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
