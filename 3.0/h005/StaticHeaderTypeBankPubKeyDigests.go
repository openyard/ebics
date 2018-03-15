// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

// complex type
type StaticHeaderTypeBankPubKeyDigests struct {
	Authentication *StaticHeaderTypeBankPubKeyDigestsAuthentication `xml:"Authentication"`
	Encryption     *StaticHeaderTypeBankPubKeyDigestsEncryption     `xml:"Encryption"`
	Signature      *StaticHeaderTypeBankPubKeyDigestsSignature      `xml:"Signature,omitempty"`
}
