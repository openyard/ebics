// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type EbicsNoPubKeyDigestsRequest struct {
	Header        EbicsNoPubKeyDigestsRequestHeader `xml:"header"`
	Body          EbicsNoPubKeyDigestsRequestBody   `xml:"body"`
	AuthSignature AuthSignature                     `xml:"AuthSignature"`
	VersionAttrGroup
}
