// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexElement
type StaticHeaderTypeBankPubKeyDigestsAuthentication struct {
	Value   *PubKeyDigestType          `xml:",chardata"`
	Version *AuthenticationVersionType `xml:"Version,attr"`
}

func NewStaticHeaderTypeBankPubKeyDigestsAuthentication() *StaticHeaderTypeBankPubKeyDigestsAuthentication {
	return new(StaticHeaderTypeBankPubKeyDigestsAuthentication)
}

func (me *StaticHeaderTypeBankPubKeyDigestsAuthentication) SetVersion(value *AuthenticationVersionType) {
	me.Version = value
}
