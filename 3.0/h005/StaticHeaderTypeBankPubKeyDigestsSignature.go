// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

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
