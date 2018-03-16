// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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
