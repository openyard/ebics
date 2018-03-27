// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexElement
type StaticHeaderTypeBankPubKeyDigestsSignature struct {
	Value   *PubKeyDigestType     `xml:",chardata"`
	Version *SignatureVersionType `xml:"Version,attr"`
}

func NewStaticHeaderTypeBankPubKeyDigestsSignature() *StaticHeaderTypeBankPubKeyDigestsSignature {
	return new(StaticHeaderTypeBankPubKeyDigestsSignature)
}

func (me *StaticHeaderTypeBankPubKeyDigestsSignature) SetVersion(value *SignatureVersionType) {
	me.Version = value
}
