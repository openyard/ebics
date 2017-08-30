// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

// complex type
type EbicsNoPubKeyDigestsRequest struct {
	Header        EbicsNoPubKeyDigestsRequestHeader `xml:"header"`
	Body          EbicsNoPubKeyDigestsRequestBody   `xml:"body"`
	AuthSignature AuthSignature                     `xml:"AuthSignature"`
	VersionAttrGroup
}
