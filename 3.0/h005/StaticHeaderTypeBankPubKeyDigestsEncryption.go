// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexElement
type StaticHeaderTypeBankPubKeyDigestsEncryption struct {
	Value   *PubKeyDigestType      `xml:",chardata"`
	Version *EncryptionVersionType `xml:"Version,attr"`
}

func NewStaticHeaderTypeBankPubKeyDigestsEncryption() *StaticHeaderTypeBankPubKeyDigestsEncryption {
	return new(StaticHeaderTypeBankPubKeyDigestsEncryption)
}

func (me *StaticHeaderTypeBankPubKeyDigestsEncryption) SetVersion(value *EncryptionVersionType) {
	me.Version = value
}
