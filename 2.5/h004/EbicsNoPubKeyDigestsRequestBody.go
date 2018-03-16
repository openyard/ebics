// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
