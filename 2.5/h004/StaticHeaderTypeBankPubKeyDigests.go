// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexType
type StaticHeaderTypeBankPubKeyDigests struct {
	Authentication *StaticHeaderTypeBankPubKeyDigestsAuthentication `xml:"Authentication"`
	Encryption     *StaticHeaderTypeBankPubKeyDigestsEncryption     `xml:"Encryption"`
	Signature      *StaticHeaderTypeBankPubKeyDigestsSignature      `xml:"Signature,omitempty"`
}

func NewStaticHeaderTypeBankPubKeyDigests() *StaticHeaderTypeBankPubKeyDigests {
	return new(StaticHeaderTypeBankPubKeyDigests)
}

func (me *StaticHeaderTypeBankPubKeyDigests) SetAuthentication(value *StaticHeaderTypeBankPubKeyDigestsAuthentication) {
	me.Authentication = value
}

func (me *StaticHeaderTypeBankPubKeyDigests) AddAuthentication() *StaticHeaderTypeBankPubKeyDigestsAuthentication {
	me.Authentication = new(StaticHeaderTypeBankPubKeyDigestsAuthentication)
	return me.Authentication
}

func (me *StaticHeaderTypeBankPubKeyDigests) SetEncryption(value *StaticHeaderTypeBankPubKeyDigestsEncryption) {
	me.Encryption = value
}

func (me *StaticHeaderTypeBankPubKeyDigests) AddEncryption() *StaticHeaderTypeBankPubKeyDigestsEncryption {
	me.Encryption = new(StaticHeaderTypeBankPubKeyDigestsEncryption)
	return me.Encryption
}

func (me *StaticHeaderTypeBankPubKeyDigests) SetSignature(value *StaticHeaderTypeBankPubKeyDigestsSignature) {
	me.Signature = value
}

func (me *StaticHeaderTypeBankPubKeyDigests) AddSignature() *StaticHeaderTypeBankPubKeyDigestsSignature {
	me.Signature = new(StaticHeaderTypeBankPubKeyDigestsSignature)
	return me.Signature
}
